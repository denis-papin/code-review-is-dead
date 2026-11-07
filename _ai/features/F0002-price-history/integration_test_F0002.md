# Integration tests for F0002

**Status:** Implemented. See the [feature tests](../../../marketdata/history_test.go) and [shared HTTP tests](../../../marketdata/http_test.go). **Boundary:** Call the public [history route](F0002.md) over HTTP with a local Yahoo fixture server. Validate status, JSON media type and response schema against [OpenAPI](../../openapi.json). For failures, assert the Error schema. Message wording is unspecified.

Use a Yahoo chart result with symbol ABC, USD currency, a timestamp array and a matching `indicators.quote[0].close` array unless the case overrides them. Call `GET /v1/history/ABC` unless another path is specified.

| ID | Fixture and action | Assertions |
| --- | --- | --- |
| H1 | Call `/v1/history/abc` with one USD daily bar. | Upstream receives chart symbol ABC, range 1mo and interval 1d. HTTP 200 identifies ABC and USD and includes the fixture sample. |
| H2 | Use November 4, 2 and 3, 2026 at 14:30 UTC, in that order. Supply closes `[103,101,null]` and distinct adjusted closes `[93,91,92]`. | HTTP 200 returns exactly two samples in order: November 2 at 101 and November 4 at 103. Both preserve their original 14:30 UTC timestamps. |
| H3 | In separate runs supply paired empty arrays and two valid timestamps with two null closes. | Each returns 200 with ABC, USD and `"samples":[]`. |
| H4 | Independently request decoded symbols `BAD!` and `ABCDEFGHIJKLMNOPQRSTU` (21 letters). Also try the 20-letter prefix with a valid chart fixture. | Each invalid request returns 400 with zero upstream calls. The 20-letter symbol reaches the provider and returns 200. |
| H5 | Return Yahoo’s explicit unknown-symbol chart error for MISSING. Call `/v1/history/MISSING`. | HTTP 404. |
| H6 | Independently omit currency, timestamp array or close array, then supply two timestamps with one close. | Each variant returns 502. |
| E1 | In separate runs return upstream 429 and hold the fixture response beyond the configured provider deadline. | Both requests return 503. |
| E2 | In separate runs return upstream 200 with malformed JSON and upstream 500. | Both requests return 502. |

## F1 · Discovery to history

Register a screener fixture returning BRK-B and a USD chart fixture for BRK-B with one unadjusted close of 101 at `2026-11-02T14:30:00Z`. Call the [F0001 route](../F0001-us-top-gainers/F0001.md), assert 200 and extract the returned symbol. Build the history request from that actual value. Assert the chart request uses BRK-B and the public response is 200 with symbol BRK-B, currency USD and exactly the fixture sample. This is one shared scenario across both features.
