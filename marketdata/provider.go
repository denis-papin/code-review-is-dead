package marketdata

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

type provider struct {
	base   string
	client *http.Client
}

func upstreamFailure(err error) int {
	if os.IsTimeout(err) || errors.Is(err, context.DeadlineExceeded) {
		return http.StatusServiceUnavailable
	}
	return http.StatusBadGateway
}

// A chart 404 may contain an explicit unknown-symbol error. Its caller checks
// the decoded envelope before exposing that status to the client.
func (p provider) fetch(ctx context.Context, path string, dest any) int {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.base+path, nil)
	if err != nil {
		return http.StatusBadGateway
	}
	req.Header.Set("User-Agent", "MarketData/1.0")
	resp, err := p.client.Do(req)
	if err != nil {
		return upstreamFailure(err)
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusTooManyRequests:
		return http.StatusServiceUnavailable
	case http.StatusOK, http.StatusNotFound:
	default:
		return http.StatusBadGateway
	}
	const limit = 2 << 20
	data, err := io.ReadAll(io.LimitReader(resp.Body, limit+1))
	if err != nil {
		return upstreamFailure(err)
	}
	if len(data) > limit || json.Unmarshal(data, dest) != nil {
		return http.StatusBadGateway
	}
	return resp.StatusCode
}
