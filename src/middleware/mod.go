package middleware

import (
	"mime"
	"os"
	"path/filepath"

	"github.com/evex-dev/hono.go/src/context"
	"github.com/evex-dev/hono.go/src/server"
)

func ServeStatic(path string) server.HandlerFunc {
	return func(ctx *context.Context) {
		ctype := CheckType(path)
		file, err := os.ReadFile(path)
		if err != nil {
			ctx.Status(404)
			ctx.Text("404 not found")
			ctx.End()
			return
		}
		ctx.Status(200)
		ctx.Body(file)
		ctx.AddHeader("Content-Type", ctype)
		ctx.End()
	}
}

func CheckType(path string) string {
	ex := filepath.Ext(path)
	return mime.TypeByExtension(ex)
}

func PoweredBy() server.HandlerFunc {
	return func(ctx *context.Context) {
		ctx.AddHeader("X-Powered-By", "Hono.go")
		ctx.Next()
	}
}
