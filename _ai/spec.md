# Market performance API

**Revision 0.2.0 · Implemented and integration-tested**

A caller can discover the five strongest percentage gainers in Yahoo Finance’s US Day Gainers screener and look up one month of daily prices for any returned symbol. The [OpenAPI contract](openapi.json) defines the HTTP interface. The [acceptance scenarios](acceptance.md) define the behavior that callers can verify.

## Top gainers

Feature [F0001 · US top gainers](features/F0001-us-top-gainers/F0001.md).

`GET /v1/markets/us/top-gainers` reads the first page of Yahoo’s predefined US Day Gainers screener with a count of five. Yahoo defines the eligible equities and applies its own price, volume and other filters. Its published ranking uses percentage change. [Yahoo Day Gainers feed](https://query1.finance.yahoo.com/v1/finance/screener/predefined/saved?scrIds=day_gainers&count=5&region=US)

Return up to five rows ordered by descending unrounded `regularMarketChangePercent`. Preserve Yahoo’s order for equal gains. A value of 5 means a 5% gain relative to the previous regular-session close. Map each row’s `regularMarketTime` to a UTC timestamp. On weekends and holidays this retains the time of the last reported trade. The endpoint describes the latest available snapshot, including a session in progress.

A valid empty list returns 200. A missing list, duplicate symbol, nonpositive gain or missing required row field makes the snapshot unusable and returns 502. Never fill missing positions with invented records.

## Daily price history

Feature [F0002 · One month of daily price history](features/F0002-price-history/F0002.md).

`GET /v1/history/{symbol}` requests Yahoo chart data for exactly one month at daily intervals (`range=1mo&interval=1d`). Validate the decoded path symbol against the OpenAPI pattern, then uppercase it. Return the currency from Yahoo and pair each timestamp with the unadjusted close at the same array index. Skip null closes, order the remaining samples oldest first and preserve each provider timestamp in UTC. The current day’s bar can change during trading. [Yahoo chart feed](https://query1.finance.yahoo.com/v8/finance/chart/AAPL?range=1mo&interval=1d)

A known symbol with paired empty arrays or all-null closes returns an empty list. An explicit unknown-symbol response returns 404. Missing currency, missing arrays or unequal array lengths return 502. Both endpoints return 503 for a Yahoo rate limit or timeout and 502 for other upstream failures.

## Module boundary and verification

One Go server registers the two routes. One marketdata package owns their business rules and Yahoo HTTP calls. Provider payloads stay private to that package. The server can import marketdata. Marketdata has no dependency on the server. Both reads can run independently.

For implementation review, exercise the public routes against a local Yahoo fixture server. Record the scenario results, `go test ./...` and `go vet ./...` outcomes. Generate the project package-import graph from `go list -json ./...` and check the allowed dependency direction. An engineer reviews scope and boundaries before implementation, then reviews code, tests and execution evidence before integration. Record accepted behavior changes in these versioned specs first. [i-smile delivery method](https://i-smile.one/method.html)
