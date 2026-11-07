package marketdata_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

const friday int64 = 1793998800

func quote(symbol string, gain float64, stamp int64) map[string]any {
	return map[string]any{"symbol": symbol, "regularMarketChangePercent": gain, "regularMarketTime": stamp}
}
func screener(rows ...map[string]any) string {
	data, err := json.Marshal(rows)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf(`{"finance":{"result":[{"quotes":%s}]}}`, data)
}
func gainer(symbol string, gain float64, stamp int64) string {
	return fmt.Sprintf(`{"symbol":%q,"changePercent":%v,"asOf":%q}`, symbol, gain, time.Unix(stamp, 0).UTC().Format(time.RFC3339))
}
func gainers(rows ...string) string { return `{"gainers":[` + strings.Join(rows, ",") + `]}` }

func TestG1FirstPage(t *testing.T) {
	rows := []map[string]any{}
	expected := []string{}
	for i, gain := range []float64{12, 10, 8, 7, 5, 4} {
		symbol := string(rune('A' + i))
		rows = append(rows, quote(symbol, gain, friday))
		if i < 5 {
			expected = append(expected, gainer(symbol, gain, friday))
		}
	}
	rows[1]["regularMarketChange"] = 1
	rows[4]["regularMarketChange"] = 5
	api := fixture(t, func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if r.Method != "GET" || r.URL.Path != "/v1/finance/screener/predefined/saved" || q.Get("region") != "US" || q.Get("scrIds") != "day_gainers" || q.Get("count") != "5" || q.Get("start") != "0" {
			t.Errorf("unexpected provider request %s %s", r.Method, r.URL)
			w.WriteHeader(400)
			return
		}
		count, _ := strconv.Atoi(q.Get("count"))
		fmt.Fprint(w, screener(rows[:count]...))
	}, 2*time.Second)
	check(t, api+gainPath, 200, gainers(expected...))
}
func TestGainers(t *testing.T) {
	for _, tc := range []struct{ name, body, want string }{
		{"G2_stable_ties", screener(quote("ZED", 7.001, friday), quote("ABC", 7.001, friday), quote("MID", 7.004, friday)), gainers(gainer("MID", 7.004, friday), gainer("ZED", 7.001, friday), gainer("ABC", 7.001, friday))},
		{"G3_last_market_time", screener(quote("A", 8, friday-60), quote("B", 7, friday)), `{"gainers":[{"symbol":"A","changePercent":8,"asOf":"2026-11-06T20:59:00Z"},{"symbol":"B","changePercent":7,"asOf":"2026-11-06T21:00:00Z"}]}`},
		{"G4_empty", `{"finance":{"result":[{"quotes":[]}]}}`, gainers()},
	} {
		t.Run(tc.name, func(t *testing.T) {
			api := fixture(t, fixed(tc.body), 2*time.Second)
			check(t, api+gainPath, 200, tc.want)
		})
	}
}
func TestG5UnusableSnapshot(t *testing.T) {
	cases := map[string]string{"missing_list": `{"finance":{"result":[{}]}}`, "duplicate": screener(quote("ABC", 7, friday), quote("ABC", 8, friday)), "zero": screener(quote("ABC", 0, friday)), "negative": screener(quote("ABC", -1, friday))}
	for _, field := range []string{"symbol", "regularMarketChangePercent", "regularMarketTime"} {
		row := quote("ABC", 7, friday)
		delete(row, field)
		cases[field] = screener(row)
	}
	for name, body := range cases {
		t.Run(name, func(t *testing.T) { api := fixture(t, fixed(body), 2*time.Second); check(t, api+gainPath, 502, "") })
	}
}
