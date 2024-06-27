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
	e.MatchRouter = NewRouter(&e.Routes)
	
	return http.ListenAndServe(addr, e)
}

func (e *Engine) RunTLS(addr string, cert, key string) error {
	e.MatchRouter = NewRouter(&e.Routes)

	return http.ListenAndServeTLS(addr, cert, key, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	matchedRoute := e.MatchRouter.Match(r.URL.Path, r.Method)

	if matchedRoute == nil {
		if e.NotFoundHandler != nil {
			e.NotFoundHandler(&context.Context{
				Writer:  w,
				Request: r,
				Params:  context.Params{},
			})
			return
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
		return
	}

	e.Method(matchedRoute[0], w, r, context.Params{})
}

func (e *Engine) Method(route *Route, w http.ResponseWriter, r *http.Request, params context.Params) {
	ctx := &context.Context{
		Writer:  w,
		Request: r,
		Params:  params,
	}

	route.Handler(ctx)
}
