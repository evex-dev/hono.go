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
			}else {
				isEnd = true
			}

			if r.IsMiddleware {
				c.Next = func() {
					m := RequestHandlerManager{
						Routes: routes,
						IsEnd: isEnd,
					}

					m.RequestHandler(c)

					isEnd = m.IsEnd
					routes = m.Routes
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
				m := RequestHandlerManager{
					Routes: routes,
					IsEnd: isEnd,
				}

				m.RequestHandler(c)

				isEnd = m.IsEnd
				routes = m.Routes
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
	Routes []*Route
}

func (m *RequestHandlerManager) RequestHandler(c *context.Context) {
	r := m.Routes[0]

	if len(m.Routes) > 1 {
		m.Routes = m.Routes[1:]
	}else {
		m.IsEnd = true
	}

	if r.IsMiddleware {
		c.Next = func() {
			m.RequestHandler(c)
		}
	} else {
		c.Next = func() {
			fmt.Println("[WARN] c.Next is only for middleware")
		}
	}

	c.End = func() {
		m.IsEnd = true
	}

	r.Handler(c)

	if len(m.Routes) == 0 {
		m.IsEnd = true
	}

	if m.IsEnd {
		return
	} else {
		m.RequestHandler(c)
	}
}
