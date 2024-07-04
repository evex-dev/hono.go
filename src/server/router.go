package server

import (
	"strings"

	"github.com/evex-dev/hono.go/src/context"
)

type Router struct {
	Match func(pattern string, method string) (FoundRoutes, *context.Params)
}

type FoundRoutes []*Route

const ALL_METHODS = "ALL"

// PathFixer a => /a
func PathFixer(path string) string {
	path = strings.Trim(path, " ")

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return path
}
