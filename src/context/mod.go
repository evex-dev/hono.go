package context

import (
	"net/http"
)

type Params map[string]string

type Context struct {
	Res    http.ResponseWriter
	Req    *http.Request
	Params Params
	End    func()
	Next   func()
	Var    map[string]any
	NotFound func()
}
