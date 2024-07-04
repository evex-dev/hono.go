package server

import (
	"fmt"

	"github.com/evex-dev/hono.go/src/context"
)

func Compose(routes ...*Route) HandlerFunc {

	sortRoutes(routes)

	return func(c *context.Context) {
		isEnd := false
		handler := func() {
			r := routes[0]

			if len(routes) > 1 {
				routes = routes[1:]
			}

			if r.IsMiddleware {
				c.Next = func() {
					rhm := RequestHandlerManager{
						IsEnd: isEnd,
					}

					rhm.RequestHandler(routes, r, c)

					isEnd = rhm.IsEnd
				}
			} else {
				c.Next = func() {
					fmt.Println("[WARN] c.Next is only for middleware")
				}
			}

			c.End = func() {
				isEnd = true
			}

			r.Handler(c)

			if len(routes) == 0 {
				isEnd = true
			}

			if isEnd {
				return
			} else {
				rhm := RequestHandlerManager{
					IsEnd: isEnd,
				}

				rhm.RequestHandler(routes, r, c)

				isEnd = rhm.IsEnd
			}
		}

		handler()
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

type RequestHandlerManager struct {
	IsEnd bool
}

func (rhm *RequestHandlerManager) RequestHandler(routes []*Route, r *Route, c *context.Context) {
	if len(routes) > 1 {
		routes = routes[1:]
	}

	if r.IsMiddleware {
		c.Next = func() {
			rhm.RequestHandler(routes, r, c)
		}
	} else {
		c.Next = func() {
			fmt.Println("[WARN] c.Next is only for middleware")
		}
	}

	c.End = func() {
		rhm.IsEnd = true
	}

	r.Handler(c)

	if len(routes) == 0 {
		rhm.IsEnd = true
	}

	if rhm.IsEnd {
		return
	} else {
		rhm.RequestHandler(routes, r, c)
	}
}
