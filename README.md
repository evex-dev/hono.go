# hono.go
Web framework based on Honojs in Golang.  
Faster x1.5 than gin.

<img src="/.github/assets/hero.png" height="200" alt="Hero" />

## Installation
```sh
go get -u github.com/evex-dev/hono.go
```

Documentation is coming soon.

## Example

### Example 1 - Minimal

```go
package main

import (
	"github.com/evex-dev/hono.go/src/context"
	"github.com/evex-dev/hono.go/src/server"
)

func main() {
	app := server.Create()

	app.GET("/", func(c *context.Context) {
		c.Status(200)
		c.WriteString("Hello World")
	})

	app.Init().Fire()
}
```

### Example 2 - Options

```go
package main

import (
	"fmt"

	"github.com/evex-dev/hono.go/src/context"
	"github.com/evex-dev/hono.go/src/server"
)

func main() {
	app := server.Create()

	app.Use("/*", func(c *context.Context) {
		fmt.Println("Catch Request on", c.URL().Path)
	})

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
```

### Example 3 - Middleware

```go
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
```