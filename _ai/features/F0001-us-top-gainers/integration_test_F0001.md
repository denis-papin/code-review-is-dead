# Integration tests for F0001

**Status:** Implemented. See the [feature tests](../../../marketdata/gainers_test.go) and [shared HTTP tests](../../../marketdata/http_test.go). **Boundary:** Call the public [top-gainers route](F0001.md) over HTTP with a local Yahoo fixture server. Validate status, JSON media type and response schema against [OpenAPI](../../openapi.json). For failures, assert the Error schema. Message wording is unspecified.

Use Yahoo’s screener envelope with `finance.result[0].quotes`. Unless a case overrides them, rows have distinct valid symbols, positive `regularMarketChangePercent` values and a valid Unix-seconds `regularMarketTime`. Each action below is one public GET to `/v1/markets/us/top-gainers`, except F1.

| ID | Fixture | Assertions |
| --- | --- | --- |
| G1 | The fixture’s six eligible symbols A through F gained 12%, 10%, 8%, 7%, 5% and 4%. B gained $1 and E gained $5. The fixture serves the requested page size. | Upstream selects the US `day_gainers` screener, first page, count 5. HTTP 200 returns A through E in that order with the five stated percentages. |
| G2 | Supply rows ZED 7.001%, ABC 7.001%, MID 7.004% in that order. | HTTP 200 returns MID, ZED, ABC with original percentages. This detects premature rounding and alphabetical tie-breaking. |
| G3 | Fix request time to Saturday, November 7, 2026. Supply only two Friday trades at 20:59 and 21:00 UTC. | HTTP 200 returns exactly two rows. Their respective timestamps are `2026-11-06T20:59:00Z` and `2026-11-06T21:00:00Z`. |
| G4 | Supply an empty quotes array. | HTTP 200 with `{"gainers":[]}`. |
| G5 | Independently omit the quotes list, repeat a symbol, set a gain to zero or negative, or remove one required row field at a time (symbol, gain, timestamp). | Each variant returns 502 for the whole snapshot. |
| E1 | In separate runs return upstream 429 and hold the fixture response beyond the configured provider deadline. | Both requests return 503. |
| E2 | In separate runs return upstream 200 with malformed JSON and upstream 500. | Both requests return 502. |
| F1 | Return BRK-B in this route’s fixture and register a chart response for it. | Execute the [shared two-request scenario](../F0002-price-history/integration_test_F0002.md#f1--discovery-to-history). |
