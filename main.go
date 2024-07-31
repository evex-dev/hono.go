package main

import (
	"fmt"

	"github.com/evex-dev/hono.go/src/context"
	"github.com/evex-dev/hono.go/src/server"
)

func main() {
	app := server.CreateHonoGo()

	app.Use("/*", func(c *context.Context) {
		fmt.Println("Catch Request on", c.URL().Path)
	})

	app.Get("/", func(c *context.Context) {
		c.Status(200)
		c.Text("Hello World")
		c.End()
	}).Get("/2", func(c *context.Context) {
		c.Status(200)
		c.Html("<b>Hello World 2</b>")
		c.End()
	}).Post("/3", func(c *context.Context) {
		c.Status(200)
		c.Body([]byte("Hello World 3"))
		c.End()
	})

	app.Init().SetHost("localhost").SetPort("3000").Callback(func(addr string, err error) error {
		fmt.Printf("Listening on http://%s\n", addr)
		return err
	}).Fire()
}
