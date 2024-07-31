package server

import (
	"net/http"

	"github.com/evex-dev/hono.go/src/context"
)

// discord look

func (e *Engine) AddRoute(method, pattern string, handler HandlerFunc, isMiddleware bool) {
	e.Routes.RouteList = append(e.Routes.RouteList, &Route{
		Method:  method,
		Pattern: pattern,
		Handler: handler,
		Index:   len(e.Routes.RouteList),
		IsMiddleware: isMiddleware,
	})
}

func (e *Engine) Run(addr string) error {
	e.MatchRouter = NewTrieRouter(&e.Routes)
	
	return http.ListenAndServe(addr, e)
}

func (e *Engine) RunTLS(addr string, cert, key string) error {
	e.MatchRouter = NewTrieRouter(&e.Routes)

	return http.ListenAndServeTLS(addr, cert, key, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	matchedRoutes, params := e.MatchRouter.Match(r.URL.Path, r.Method)

	notFoundHandler := func(c *context.Context) {
		if e.NotFoundHandler != nil {
			e.NotFoundHandler(c)
			return
		}
		c.Status(404)
		c.Text("404 Not Found")
		c.End()
	}

	e.Serve(matchedRoutes, w, r, *params, notFoundHandler)
}

func (e *Engine) Serve(routes []*Route, w http.ResponseWriter, r *http.Request, params context.Params, notFoundHandler HandlerFunc) {
	ctx := &context.Context{
		Res:  w,
		Req: r,
		Params:  params,
	}

	if countHandler(routes) == 0 {
		routes = append(routes, &Route{
			Handler: notFoundHandler,
			Index:   0,
			IsMiddleware: false,
		})
	}

	Compose(routes...)(ctx)
}

func countHandler(routes []*Route) int {
	count := 0
	for _, route := range routes {
		if !route.IsMiddleware {
			count++
		}
	}
	return count
}
