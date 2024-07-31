package context

import (
	"encoding/json"
)

func (ctx *Context) Body(value []byte) *Context {
	ctx.Res.Write(value)
	return ctx
}
func (ctx *Context) Text(value string) *Context {
	ctx.SetHeader("Content-Type", "text/plain; charset=utf-8")
	ctx.Body([]byte(value))
	return ctx
}

func (ctx *Context) Json(value any) *Context {
	marshal, _ := json.Marshal(value)
	ctx.SetHeader("Content-Type", "application/json; charset=utf-8")
	ctx.Body(marshal)
	return ctx
}

func (ctx *Context) Html(value string) *Context {
	ctx.SetHeader("Content-Type", "text/html; charset=utf-8")
	ctx.Body([]byte(value))
	return ctx
}

