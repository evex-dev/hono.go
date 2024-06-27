package main

import (
	"fmt"
	"github.com/evex-dev/hono.go/src/context"
	"github.com/evex-dev/hono.go/src/server"
)

func main() {
	app := server.CreateHonoGo()

	app.GET("/", func(c *context.Context) {
		c.Status(200)
		c.WriteString("Hello World")
	}).POST("/2", func(c *context.Context) {
		c.Status(200)
		c.Write([]byte("Hello World 2"))
	})


	app.Init().SetPort("3000").Callback(func(addr string, err error) error {
		fmt.Printf("Listening on http://localhost%s\n", addr)
		return err
	}).Fire()
}
