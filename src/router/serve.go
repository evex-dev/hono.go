package router

import "net/http"

func (e *Engine) AddRoute(method, pattern string, handler HandlerFunc) {
	e.Route = append(e.Route, &ContextMethod{
		Method:  method,
		Pattern: pattern,
		Handler: handler,
	})
}

func (e *Engine) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	for _,r:=range e.Route{
		
	}
}
