package context

import (
	"encoding/json"
)

func (c *Context) Body(value []byte) *Context {
	c.Res.Write(value)
	return c
}
func (c *Context) Text(value string) *Context {
	c.SetHeader("Content-Type", "text/plain; charset=utf-8")
	c.Body([]byte(value))
	return c
}

func (c *Context) Json(value any) *Context {
	marshal, _ := json.Marshal(value)
	c.SetHeader("Content-Type", "application/json; charset=utf-8")
	c.Body(marshal)
	return c
}

func (c *Context) Html(value string) *Context {
	c.SetHeader("Content-Type", "text/html; charset=utf-8")
	c.Body([]byte(value))
	return c
}
