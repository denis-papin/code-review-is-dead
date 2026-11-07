package marketdata

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"
	"time"
)

type quote struct {
	Symbol    string
	Gain      *float64 `json:"regularMarketChangePercent"`
	Timestamp *int64   `json:"regularMarketTime"`
}
type gainer struct {
	Symbol        string    `json:"symbol"`
	ChangePercent float64   `json:"changePercent"`
	AsOf          time.Time `json:"asOf"`
}

func (p provider) topGainers(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Finance struct {
			Error  *json.RawMessage
			Result []struct{ Quotes []quote }
		}
	}
	status := p.fetch(r.Context(), "/v1/finance/screener/predefined/saved?scrIds=day_gainers&region=US&start=0&count=5", &payload)
	if status != http.StatusOK {
		if status == http.StatusNotFound {
			status = http.StatusBadGateway
		}
		fail(w, status)
		return
	}
	result := payload.Finance.Result
	if payload.Finance.Error != nil || len(result) != 1 || result[0].Quotes == nil {
		fail(w, http.StatusBadGateway)
		return
	}
	gainers := make([]gainer, 0, len(result[0].Quotes))
	seen := make(map[string]bool)
	for _, q := range result[0].Quotes {
		key := strings.ToUpper(q.Symbol)
		if !symbolPattern.MatchString(q.Symbol) || seen[key] || q.Gain == nil || *q.Gain <= 0 || q.Timestamp == nil {
			fail(w, http.StatusBadGateway)
			return
		}
		seen[key] = true
		gainers = append(gainers, gainer{q.Symbol, *q.Gain, utc(*q.Timestamp)})
	}
	sort.SliceStable(gainers, func(i, j int) bool { return gainers[i].ChangePercent > gainers[j].ChangePercent })
	send(w, http.StatusOK, map[string][]gainer{"gainers": gainers[:min(len(gainers), 5)]})
}
