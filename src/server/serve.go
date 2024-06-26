package server

import (
	"net/http"

	"github.com/evex-dev/hono.go/src/context"
)

func (e *Engine) AddRoute(method, pattern string, handler HandlerFunc) {
	e.Routes = append(e.Routes, &Route{
		Method:  method,
		Pattern: pattern,
		Handler: handler,
	})
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) RunTLS(addr string, cert, key string) error {
	return http.ListenAndServeTLS(addr, cert, key, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range e.Routes {
		e.Method(route, w, r)
	}
}

func (e *Engine) Method(route *Route, w http.ResponseWriter, r *http.Request) {
	ctx := &context.Context{
		Writer:  w,
		Request: r,
		Params:  context.Params{},
	}

	route.Handler(ctx)
}
