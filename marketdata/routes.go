package marketdata

import (
	"net/http"
)

// Handler serves the market data API using the supplied Yahoo origin and client.
func Handler(base string, client *http.Client) http.Handler {
	p := provider{base, client}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/markets/us/top-gainers", p.topGainers)
	mux.HandleFunc("GET /v1/history/{symbol}", p.history)
	return mux
}
