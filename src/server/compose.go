package server

import (
	"fmt"

	"github.com/evex-dev/hono.go/src/context"
)

func Compose(routes ...*Route) HandlerFunc {

	sortedRotues := sortRoutes(routes)

	for i := 0; i < len(sortedRotues); i++ {
		fmt.Println(sortedRotues[i].Method, sortedRotues[i].Pattern)
	}

	return func(c *context.Context) {
		handler := func() {
			m := RequestHandlerManager{
				Routes: sortedRotues,
				IsEnd:  false,
			}
			m.RequestHandler(c)
		}

		handler()
	}
}

func sortRoutes(routes []*Route) []*Route {

	for i := 0; i < len(routes); i++ {
		for j := 0; j < len(routes); j++ {
			if routes[i].IsMiddleware && !routes[j].IsMiddleware {
				routes[i], routes[j] = routes[j], routes[i]
			} else if !routes[i].IsMiddleware && routes[j].IsMiddleware {
				routes[i], routes[j] = routes[j], routes[i]
			} else {
				if routes[i].Index > routes[j].Index {
					routes[i], routes[j] = routes[j], routes[i]
				}
			}
		}
	}

	return routes
}

type RequestHandlerManager struct {
	IsEnd  bool
	Routes []*Route
}

func (m *RequestHandlerManager) RequestHandler(c *context.Context) {
	r := m.Routes[0]

	if len(m.Routes) > 0 {
		m.Routes = m.Routes[1:]
	}

	c.Next = func() {
		m.RequestHandler(c)
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
