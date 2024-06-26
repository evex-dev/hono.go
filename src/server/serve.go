package server

import (
	"net/http"
	"github.com/evex-dev/hono.go/src/context"
)
// discord look

func (e *Engine) AddRoute(method, pattern string, handler HandlerFunc) {
	e.Route = append(e.Route, &Route{
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
	route := NewRouter(&e.Routes)

	matchedRoute := route.Match(r.URL.Path)

	if matchedRoute == nil {
		w.WriteHeader(http.StatusNotFound) // 後でNotFoundHandlerに変更する
		w.Write([]byte("Not Found"))
		return
	}

	e.Method(matchedRoute, w, r, context.Params{})
}

func (e *Engine) Method(route *Route, w http.ResponseWriter, r *http.Request, params context.Params) {
	ctx := &context.Context{
		Writer:  w,
		Request: r,
		Params:  params,
	}

	route.Handler(ctx)
}
