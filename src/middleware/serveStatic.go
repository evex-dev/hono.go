package middleware

import (
	"mime"
	"os"
	"path/filepath"

	"github.com/evex-dev/hono.go/src/context"
	"github.com/evex-dev/hono.go/src/server"
)

func ServeStatic(path string) server.HandlerFunc {

	rpath := []rune(path)
	if rpath[len(rpath)-1] == '/' {
		path = string(rpath[:len(rpath)-2])
	}

	return func(c *context.Context) {
		ctype := checkType(c.URL().Path)
		file, err := os.ReadFile(path + c.URL().RawPath)
		if err != nil {
			c.NotFound()
			return
		}
		c.Status(200)
		c.Body(file)
		c.SetHeader("Content-Type", ctype)
	}
}

func checkType(path string) string {
	ex := filepath.Ext(path)
	return mime.TypeByExtension(ex)
}
