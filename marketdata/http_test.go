package marketdata_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"example.com/marketdata/marketdata"
)

const gainPath = "/v1/markets/us/top-gainers"

func fixture(t *testing.T, upstream http.HandlerFunc, timeout time.Duration) string {
	t.Helper()
	yahoo := httptest.NewServer(upstream)
	t.Cleanup(yahoo.Close)
	api := httptest.NewServer(marketdata.Handler(yahoo.URL, &http.Client{Timeout: timeout}))
	t.Cleanup(api.Close)
	return api.URL
}

// Check the public status, JSON media type and complete expected JSON value.
func check(t *testing.T, endpoint string, status int, want string) map[string]any {
	t.Helper()
	resp, err := http.Get(endpoint)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != status || resp.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("got %d %s: %s", resp.StatusCode, resp.Header.Get("Content-Type"), data)
	}
	var got map[string]any
	if err = json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	assertContract(t, resp.Request.URL.Path, status, got)
	if status != 200 {
		if _, ok := got["error"].(string); !ok {
			t.Fatalf("invalid error body: %s", data)
		}
	} else {
		var expected map[string]any
		if err = json.Unmarshal([]byte(want), &expected); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("got %s, want %s", data, want)
		}
	}
	return got
}

func fixed(body string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, body) }
}

func TestProviderFailures(t *testing.T) {
	for _, path := range []string{gainPath, "/v1/history/ABC"} {
		for _, tc := range []struct {
			name         string
			status, want int
			body         string
		}{
			{"E1_rate_limit", 429, 503, ""}, {"E1_timeout", 0, 503, ""},
			{"E2_malformed", 200, 502, "{"}, {"E2_server_error", 500, 502, ""},
		} {
			t.Run(path+"/"+tc.name, func(t *testing.T) {
				deadline := 2 * time.Second
				if tc.status == 0 {
					deadline = 100 * time.Millisecond
				}
				api := fixture(t, func(w http.ResponseWriter, r *http.Request) {
					if tc.status == 0 {
						<-r.Context().Done()
						return
					}
					w.WriteHeader(tc.status)
					fmt.Fprint(w, tc.body)
				}, deadline)
				check(t, api+path, tc.want, "")
			})
		}
	}
}
