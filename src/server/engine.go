package server

import (
	"github.com/evex-dev/hono.go/src/context"
)

type HandlerFunc func(*context.Context)

type Route struct {
	Method       string
	Pattern      string
	Handler      HandlerFunc
	Index        int
	IsMiddleware bool
}

type Routes struct {
	RouteList       []*Route
	NotFoundHandler HandlerFunc
}

type Engine struct {
	Routes
	MatchRouter *Router
}
