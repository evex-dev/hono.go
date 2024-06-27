package json

import (
	"fmt"
	"testing"
)

func Parser_Test(t *testing.T) {
	type parse_test = struct {
		Name string `json:"name"`
	}
	value := `{"name": "hello!"}`
	var p parse_test
	ParseJSON(value, &p)
	fmt.Println(p.Name)
}
