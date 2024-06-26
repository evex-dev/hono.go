package server

import (
	"github.com/evex-dev/hono.go/src/context"
)

type HandlerFunc func(*context.Context)

type Route struct {
	Method  string
	Pattern string
	Handler HandlerFunc
}

type Engine struct {
	Routes []*Route
}
