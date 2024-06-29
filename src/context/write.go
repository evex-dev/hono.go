package context

import (
	"encoding/json"
)

func (ctx *Context) Body(value []byte) error {
	_, err := ctx.Res.Write(value)
	return err
}
func (ctx *Context) Text(value string) error {
	ctx.SetHeader("Content-Type", "text/plain; charset=utf-8")
	return ctx.Body([]byte(value))
}

func (ctx *Context) Json(value any) error {
	marshal, err := json.Marshal(value)
	if err != nil {
		return err
	}
	ctx.SetHeader("Content-Type", "application/json; charset=utf-8")
	return ctx.Body(marshal)
}

func (ctx *Context) Html(value string) error {
	ctx.SetHeader("Content-Type", "text/html; charset=utf-8")
	return ctx.Body([]byte(value))
}

