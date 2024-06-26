package main

import (
	"github.com/evex-dev/hono.go/src/server"
	"github.com/evex-dev/hono.go/src/context"	
)

func main() {
	app := server.NewHonoGo()

	app.GET("/", func(c *context.Context) {
		c.Writer.Write([]byte("Hello, World!"))
	}).GET("/a", func(c *context.Context) {
		c.Writer.Write([]byte("Hello, World!"))
	})

	app.Run(":3000")
}