package router

import (
	"github.com/evex-dev/hono.go/src/server"
)

type Router struct {
	Match func(path string) *server.Route
}

func NewRouter(routes []*server.Route) *Router {
	return &Router{
		Match: func(path string) *server.Route {
			for _, route := range routes {
				if route.Pattern == path {
					return route
				}
			}
			return nil
		},
	}
}
