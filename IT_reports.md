# Integration test report

**All 14 integration scenarios passed** on November 7, 2026. The uncached run took **3.632 seconds**, with no failures or skips. Each scenario called the Go REST API over HTTP with a local fixture server supplying Yahoo’s responses.

Run command: `go test -count=1 -json ./marketdata`

| ID | Endpoint | Verified outcome | Result |
|---|---|---|---|
| G1 | Top gainers | Requested the first five US day gainers and returned them in percentage-gain order. | Pass |
| G2 | Top gainers | Ranked 7.004% above 7.001% and preserved Yahoo’s order for equal gains. | Pass |
| G3 | Top gainers | Returned both available rows with their original Friday trade timestamps. | Pass |
| G4 | Top gainers | An empty screener returned `200` with an empty list. | Pass |
| G5 | Top gainers | All seven malformed snapshot variants returned `502`, including duplicate symbols, nonpositive gains, and missing fields. | Pass |
| H1 | History | A request for `abc` queried `ABC` for one month of daily prices and returned USD. | Pass |
| H2 | History | Returned unadjusted closes oldest first, skipped a null close, and preserved the timestamps. | Pass |
| H3 | History | Empty arrays and all-null closes both returned `200` with an empty sample list. | Pass |
| H4 | History | Invalid symbols returned `400` before any Yahoo call. A valid 20-character symbol succeeded. | Pass |
| H5 | History | An explicit unknown-symbol response from Yahoo returned `404`. | Pass |
| H6 | History | Missing currency, missing timestamps, missing closes, and unequal array lengths each returned `502`. | Pass |
| E1 | Both | A Yahoo rate limit and a timeout each returned `503` on both endpoints. | Pass |
| E2 | Both | Malformed Yahoo JSON and an upstream server error each returned `502` on both endpoints. | Pass |
| F1 | Discovery → history | Used `BRK-B` from the gainers response to fetch its history and verified the expected USD price sample. | Pass |
