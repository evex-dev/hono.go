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

type Routes struct{
	Route []*Route
}

type Engine struct {
	Routes
}

type Router struct {
	Match func(path string) *Route
}
