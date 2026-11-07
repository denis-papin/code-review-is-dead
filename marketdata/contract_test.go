package marketdata_test

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

// Select the schema by operation and actual status, so changes to the public
// contract are checked by every scenario that exercises the HTTP handler.
func assertContract(t *testing.T, route string, status int, body any) {
	t.Helper()
	const resource = "https://contract.invalid/openapi.json"
	data, err := os.ReadFile("../_ai/openapi.json")
	if err != nil {
		t.Fatal(err)
	}
	var doc map[string]any
	if err = json.Unmarshal(data, &doc); err != nil {
		t.Fatal(err)
	}
	compiler := jsonschema.NewCompiler()
	compiler.AssertFormat()
	if err = compiler.AddResource(resource, doc); err != nil {
		t.Fatal(err)
	}
	op := route
	if route != gainPath {
		op = "/v1/history/{symbol}"
	}
	pointer := "#/paths/" + strings.ReplaceAll(op, "/", "~1") + "/get/responses/" + strconv.Itoa(status) + "/content/application~1json/schema"
	schema, err := compiler.Compile(resource + pointer)
	if err != nil {
		t.Fatal(err)
	}
	if err = schema.Validate(body); err != nil {
		t.Fatal(err)
	}
}
