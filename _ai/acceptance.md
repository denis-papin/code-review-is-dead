# Acceptance scenarios

Use synthetic Yahoo responses and call the public HTTP routes. Each response must match the status and JSON schema in the [contract](openapi.json).

Detailed fixtures and assertions are specified in [F0001 integration tests](features/F0001-us-top-gainers/integration_test_F0001.md) and [F0002 integration tests](features/F0002-price-history/integration_test_F0002.md). E1 and E2 apply to each route. F1 exercises both routes together.

| ID | Given | When | Expected |
| --- | --- | --- | --- |
| G1 | Yahoo ranks six eligible equities by gains of 12%, 10%, 8%, 7%, 5% and 4%. The 10% stock gained $1 and the 5% stock gained $5. | Request US top gainers. | The upstream request selects US Day Gainers, count five, first page. The response lists gains 12, 10, 8, 7 and 5 in order. |
| G2 | Two of Yahoo’s returned rows have identical percentage gains. | Request top gainers. | Their order matches the feed. |
| G3 | On Saturday the feed contains two rows timestamped Friday. | Request top gainers. | Return 200 with both Friday timestamps in UTC. |
| G4 | Yahoo returns a valid empty quotes array. | Request top gainers. | Return 200 with an empty gainers array. |
| G5 | Independently, a list is missing, a symbol repeats, a gain is zero or negative, or a row lacks its symbol, percentage or timestamp. | Request top gainers. | Each fixture returns 502. |
| H1 | The provider has USD daily bars for ABC. | Request history for abc. | Yahoo receives ABC with range 1mo and interval 1d. The response identifies ABC and USD. |
| H2 | Daily bars arrive in date order D3, D1, D2 with closes 103, 101 and null. Adjusted closes differ. | Request history. | Return only D1 at 101 and D3 at 103, with the original timestamps in UTC. |
| H3 | A known symbol has paired empty arrays or all closes are null. | Request history. | Each fixture returns 200 with an empty samples array. |
| H4 | The decoded symbol is BAD! or exceeds 20 characters. | Request history. | Return 400 before contacting Yahoo. |
| H5 | Yahoo explicitly reports an unknown symbol. | Request history. | Return 404. |
| H6 | Currency is missing, an array is missing, or the arrays have different lengths. | Request history. | Each fixture returns 502. |
| E1 | Yahoo returns HTTP 429 or times out. | Call each route against each fixture. | Return 503. |
| E2 | Yahoo returns invalid JSON or an unexpected server error. | Call each route against each fixture. | Return 502. |
| F1 | The screener returns BRK-B and the chart fixture has history for BRK-B. | Pass the returned symbol to the history route. | Both calls succeed and identify the same symbol. |
