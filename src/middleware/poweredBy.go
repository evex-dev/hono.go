package middleware

import (
	"github.com/evex-dev/hono.go/src/context"
	"github.com/evex-dev/hono.go/src/server"
)

func PoweredBy() server.HandlerFunc {
	return func(c *context.Context) {
		c.SetHeader("X-Powered-By", "Hono.go")
		c.Next()
	}
}
