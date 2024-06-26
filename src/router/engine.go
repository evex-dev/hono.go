package router

import (
	"net/http"
	"github.com/evex-dev/hono.go/context"
)

type Params map[string]string

type Context struct {
	Write   http.ResponseWriter
	Request *http.Request
	Params  Params
}

type HandlerFunc func(*Context)

type ContextMethod struct {
	Method  string
	Pattern string
	Handler HandlerFunc
}

type Engine struct {
	Route []*ContextMethod
}
