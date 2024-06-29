package context

import (
	"net/http"
)

type Params map[string]string

type Context struct {
	Res    http.ResponseWriter
	Req    *http.Request
	Params Params
}
