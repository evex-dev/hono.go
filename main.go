package main

import (
	"fmt"

	"github.com/evex-dev/hono.go/src/context"
	"github.com/evex-dev/hono.go/src/middleware"
	"github.com/evex-dev/hono.go/src/server"
)

func main() {
	app := server.Create()

	app.Use("/*", func(c *context.Context) {
		fmt.Println("Catch Request on", c.URL().Path)
	}).Use("/*", middleware.PoweredBy())

	app.Get("/", func(c *context.Context) {
		c.Status(200)
		c.Text("Hello World")
	}).Get("/2", func(c *context.Context) {
		c.Status(200).Html("<b>Hello World 2</b>")
	}).Post("/3", func(c *context.Context) {
		c.Status(200).Body([]byte("Hello World 3")).End()
	})

	app.Init().SetHost("localhost").SetPort("3000").Callback(func(addr string, err error) error {
		fmt.Printf("Listening on http://%s\n", addr)
		return err
	}).Fire()
}
