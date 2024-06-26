package server

import (
	"net/http"
	"github.com/evex-dev/hono.go/router"
)

func Run(addr string, engine *router.Engine) {
	http.ListenAndServe(":8080", nil)
}