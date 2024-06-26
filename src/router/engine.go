package router

import (
	"net/http"
	"github.com/evex-dev/hono.go/context"
)

type HandlerFunc func(*context.Context)

type ContextMethod struct {
	Method  string
	Pattern string
	Handler HandlerFunc
}

type Engine struct {
	Route []*ContextMethod
}
