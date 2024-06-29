package server

import "github.com/evex-dev/hono.go/src/context"

func Compose(routes... *Route) HandlerFunc {

	sortRoutes(routes)

	return func(c *context.Context) {
		for _, r := range routes {
			r.Handler(c)
		}
	}
}


func sortRoutes(routes []*Route) []*Route {

	for i := 0; i < len(routes); i++ {
		for j := 0; j < len(routes); j++ {
			if routes[i].IsMiddleware && !routes[j].IsMiddleware {
				routes[i], routes[j] = routes[j], routes[i]
			}

			if !routes[i].IsMiddleware && routes[j].IsMiddleware {
				routes[i], routes[j] = routes[j], routes[i]
			}

			if routes[i].Index > routes[j].Index {
				routes[i], routes[j] = routes[j], routes[i]
			}
		}
	}

	return routes
}