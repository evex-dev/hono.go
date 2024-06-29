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

	if matchedRoutes == nil {
		if e.NotFoundHandler != nil {
			e.NotFoundHandler(&context.Context{
				Res:  w,
				Req: r,
				Params:  context.Params{},
			})
			return
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
		return
	}

	e.Method(matchedRoutes, w, r, *params)
}

func (e *Engine) Method(routes []*Route, w http.ResponseWriter, r *http.Request, params context.Params) {
	ctx := &context.Context{
		Res:  w,
		Req: r,
		Params:  params,
	}

	Compose(routes...)(ctx)
}
