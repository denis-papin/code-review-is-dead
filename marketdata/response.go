package marketdata

import (
	"encoding/json"
	"net/http"
)

func send(w http.ResponseWriter, status int, body any) {
	data, err := json.Marshal(body)
	if err != nil {
		status = http.StatusBadGateway
		data = []byte(`{"error":"Unusable provider data"}`)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(data)
}

func fail(w http.ResponseWriter, status int) {
	send(w, status, map[string]string{"error": http.StatusText(status)})
}
