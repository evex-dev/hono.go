package context

import (
	"net/http"
)

type Params map[string]string

type Context struct {
	Write   http.ResponseWriter
	Request *http.Request
	Params  Params
}
