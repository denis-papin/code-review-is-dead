package main

import (
	"log"
	"net/http"
	"time"

	"example.com/marketdata/marketdata"
)

func main() {
	client := &http.Client{Timeout: 5 * time.Second}
	server := &http.Server{
		Addr:              "127.0.0.1:8080",
		Handler:           marketdata.Handler("https://query1.finance.yahoo.com", client),
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("Serving market data at http://%s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
