# Market data API

Run `go run .` from the project root with Go 1.22 or later. The server listens on `http://127.0.0.1:8080` and uses Yahoo Finance for market data.

| Request | Response |
| --- | --- |
| `GET /v1/markets/us/top-gainers` | Up to five US Day Gainers equities, ordered by session percentage gain, with each trade’s UTC timestamp. |
| `GET /v1/history/{symbol}` | One month of daily unadjusted closes in the provider’s currency, oldest first. |

To try the two endpoints together in PowerShell:

```powershell
$market = Invoke-RestMethod 'http://127.0.0.1:8080/v1/markets/us/top-gainers'
$market.gainers
if ($market.gainers.Count) {
    $symbol = $market.gainers[0].symbol
    Invoke-RestMethod "http://127.0.0.1:8080/v1/history/$symbol"
}
```

See the [OpenAPI contract](_ai/openapi.json), [F0001 specification](_ai/features/F0001-us-top-gainers/F0001.md) and [F0002 specification](_ai/features/F0002-price-history/F0002.md).

## Verification

```sh
go test -count=1 ./...
go vet ./...
```

The integration tests exercise the public routes against a local Yahoo fixture server. Each scenario checks the response against the matching OpenAPI schema and verifies the business outcomes. The runtime uses the Go standard library.

