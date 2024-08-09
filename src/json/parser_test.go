package json_test

import (
	"testing"

	"github.com/evex-dev/hono.go/src/json"
)

func TestParseJSON(t *testing.T) {
	type parse_test = struct {
		Name string `json:"name"`
	}
	value := `{"name": "hello!"}`
	var p parse_test
	json.ParseJSON(value, &p)
	t.Log(p)
}
