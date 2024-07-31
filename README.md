# hono.go
Web framework based on Honojs in Golang.

- 日本語
絶賛開発中なのでどんどんissue/PR建てて下さい。
- English
We are in the process of developing this project, so please keep on building issues/PRs.

## Example

### Example 1 - Minimal

```go
package main

import (
	"github.com/evex-dev/hono.go/src/context"
	"github.com/evex-dev/hono.go/src/server"
)

func main() {
	app := server.CreateHonoGo()

	app.GET("/", func(c *context.Context) {
		c.Status(200)
		c.WriteString("Hello World")
		c.End()
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
	app := server.CreateHonoGo()

	app.GET("/", func(c *context.Context) {
		c.Status(200)
		c.WriteString("Hello World")
	}).POST("/2", func(c *context.Context) {
		c.Status(200)
		c.Write([]byte("Hello World 2"))
	})


	app.Init().SetPort("3000").Callback(func (addr string, err error) error {
		fmt.Printf("Listening on http://localhost%s\n", addr)
		return err
	}).Fire()
}
```
