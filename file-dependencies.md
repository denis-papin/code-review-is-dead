# Yahoo API file dependencies

Read the diagram from top to bottom. Each group contains files with the same number of **direct project dependencies**. An arrow points from a file to one it uses. A file name can appear again as the target of an arrow in another group.

```mermaid
flowchart TB
  subgraph zero["0 dependencies"]
    foundation["_ai/openapi.json · provider.go · values.go · response.go"]
  end

  subgraph one["1 dependency"]
    direction LR
    api["marketdata/openapi.json"] --> schema1["_ai/openapi.json"]
    gainTest1["gainers_test.go"] --> http1["http_test.go"]
    entry["main.go"] --> routes1["routes.go"]
  end

  subgraph two["2 dependencies"]
    direction LR
    contract2["contract_test.go"] --> schema2["_ai/openapi.json"]
    contract2 --> http2["http_test.go"]
    http2 --> contract2
    http2 --> routes2["routes.go"]
    histTest["history_test.go"] --> http2
    histTest --> gainTest2["gainers_test.go"]
  end

  subgraph three["3 dependencies"]
    direction LR
    routes3["routes.go"] --> provider3["provider.go"]
    routes3 --> gain["gainers.go"]
    routes3 --> hist["history.go"]
    gain --> provider3
    gain --> values3["values.go"]
    gain --> response3["response.go"]
    hist --> provider3
    hist --> values3
    hist --> response3
  end

  zero ~~~ one
  one ~~~ two
  two ~~~ three
```

All Go files except `main.go` live in `marketdata/`. The counts include Go symbol references, the schema checker's read of the OpenAPI contract, and the module OpenAPI file's references to that contract.

The test helpers depend on each other. `http_test.go` calls `assertContract` from `contract_test.go`, and `contract_test.go` uses the `gainPath` constant from `http_test.go`.
