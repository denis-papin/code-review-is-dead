package marketdata

import (
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type chartResult struct {
	Meta       struct{ Symbol, Currency string }
	Timestamp  []*int64
	Indicators struct{ Quote []struct{ Close []*float64 } }
}
type sample struct {
	At    time.Time `json:"at"`
	Close float64   `json:"close"`
}
type historyResponse struct {
	Symbol   string   `json:"symbol"`
	Currency string   `json:"currency"`
	Samples  []sample `json:"samples"`
}

func (p provider) history(w http.ResponseWriter, r *http.Request) {
	symbol := r.PathValue("symbol")
	if !symbolPattern.MatchString(symbol) {
		fail(w, http.StatusBadRequest)
		return
	}
	symbol = strings.ToUpper(symbol)
	var payload struct {
		Chart struct {
			Result []chartResult
			Error  *struct{ Code string }
		}
	}
	status := p.fetch(r.Context(), "/v8/finance/chart/"+url.PathEscape(symbol)+"?range=1mo&interval=1d", &payload)
	if (status == http.StatusOK || status == http.StatusNotFound) && payload.Chart.Error != nil && payload.Chart.Error.Code == "Not Found" {
		fail(w, http.StatusNotFound)
		return
	}
	if status != http.StatusOK {
		if status == http.StatusNotFound {
			status = http.StatusBadGateway
		}
		fail(w, status)
		return
	}
	if payload.Chart.Error != nil || len(payload.Chart.Result) != 1 {
		fail(w, http.StatusBadGateway)
		return
	}
	chart := payload.Chart.Result[0]
	if chart.Meta.Currency == "" || chart.Timestamp == nil || len(chart.Indicators.Quote) != 1 {
		fail(w, http.StatusBadGateway)
		return
	}
	closes := chart.Indicators.Quote[0].Close
	if closes == nil || len(closes) != len(chart.Timestamp) {
		fail(w, http.StatusBadGateway)
		return
	}
	samples := make([]sample, 0, len(closes))
	for i, close := range closes {
		if chart.Timestamp[i] == nil {
			fail(w, http.StatusBadGateway)
			return
		}
		if close != nil {
			samples = append(samples, sample{utc(*chart.Timestamp[i]), *close})
		}
	}
	sort.SliceStable(samples, func(i, j int) bool { return samples[i].At.Before(samples[j].At) })
	send(w, http.StatusOK, historyResponse{symbol, chart.Meta.Currency, samples})
}
