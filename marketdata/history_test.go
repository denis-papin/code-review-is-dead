package marketdata_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func chart(symbol, stamps, closes string) string {
	return fmt.Sprintf(`{"chart":{"result":[{"meta":{"symbol":%q,"currency":"USD"},"timestamp":%s,"indicators":{"quote":[{"close":%s}],"adjclose":[{"adjclose":[93,91,92]}]}}]}}`, symbol, stamps, closes)
}

const oneSample = `[{"at":"2026-11-02T14:30:00Z","close":101}]`

func history(symbol, samples string) string {
	return fmt.Sprintf(`{"symbol":%q,"currency":"USD","samples":%s}`, symbol, samples)
}
func chartFixture(t *testing.T, symbol string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if r.Method != "GET" || r.URL.Path != "/v8/finance/chart/"+symbol || q.Get("range") != "1mo" || q.Get("interval") != "1d" {
			t.Errorf("unexpected provider request %s %s", r.Method, r.URL)
		}
		fmt.Fprint(w, chart(symbol, "[1793629800]", "[101]"))
	}
}
func TestH1NormalizedSymbol(t *testing.T) {
	api := fixture(t, chartFixture(t, "ABC"), 2*time.Second)
	check(t, api+"/v1/history/abc", 200, history("ABC", oneSample))
}
func TestH2UnadjustedCloses(t *testing.T) {
	api := fixture(t, fixed(chart("ABC", "[1793802600,1793629800,1793716200]", "[103,101,null]")), 2*time.Second)
	check(t, api+"/v1/history/ABC", 200, history("ABC", `[{"at":"2026-11-02T14:30:00Z","close":101},{"at":"2026-11-04T14:30:00Z","close":103}]`))
}
func TestH3EmptySamples(t *testing.T) {
	for _, pair := range [][2]string{{"[]", "[]"}, {"[1,2]", "[null,null]"}} {
		api := fixture(t, fixed(chart("ABC", pair[0], pair[1])), 2*time.Second)
		check(t, api+"/v1/history/ABC", 200, history("ABC", "[]"))
	}
}
func TestH4SymbolBoundaries(t *testing.T) {
	api := fixture(t, func(w http.ResponseWriter, r *http.Request) { t.Error("invalid symbol reached Yahoo") }, 2*time.Second)
	for _, bad := range []string{"BAD%21", "ABCDEFGHIJKLMNOPQRSTU"} {
		check(t, api+"/v1/history/"+bad, 400, "")
	}
	const valid = "ABCDEFGHIJKLMNOPQRST"
	api = fixture(t, chartFixture(t, valid), 2*time.Second)
	check(t, api+"/v1/history/"+valid, 200, history(valid, oneSample))
}
func TestH5UnknownSymbol(t *testing.T) {
	api := fixture(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprint(w, `{"chart":{"result":null,"error":{"code":"Not Found","description":"No data found"}}}`)
	}, 2*time.Second)
	check(t, api+"/v1/history/MISSING", 404, "")
}
func TestH6MalformedHistory(t *testing.T) {
	for name, result := range map[string]string{
		"currency":   `{"timestamp":[],"indicators":{"quote":[{"close":[]}]}}`,
		"timestamps": `{"meta":{"currency":"USD"},"indicators":{"quote":[{"close":[]}]}}`,
		"closes":     `{"meta":{"currency":"USD"},"timestamp":[],"indicators":{"quote":[{}]}}`,
		"length":     `{"meta":{"currency":"USD"},"timestamp":[1,2],"indicators":{"quote":[{"close":[1]}]}}`,
	} {
		t.Run(name, func(t *testing.T) {
			api := fixture(t, fixed(`{"chart":{"result":[`+result+`]}}`), 2*time.Second)
			check(t, api+"/v1/history/ABC", 502, "")
		})
	}
}
func TestF1DiscoveryToHistory(t *testing.T) {
	yahoo := http.NewServeMux()
	yahoo.HandleFunc("/v1/finance/screener/predefined/saved", fixed(screener(quote("BRK-B", 5.25, friday))))
	yahoo.HandleFunc("/v8/finance/chart/BRK-B", chartFixture(t, "BRK-B"))
	api := fixture(t, yahoo.ServeHTTP, 2*time.Second)
	got := check(t, api+gainPath, 200, gainers(gainer("BRK-B", 5.25, friday)))
	symbol := got["gainers"].([]any)[0].(map[string]any)["symbol"].(string)
	check(t, api+"/v1/history/"+symbol, 200, history(symbol, oneSample))
}
