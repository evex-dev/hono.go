package context

import (
	"encoding/json"
)

func (ctx *Context) BODY(value []byte) error {
	_, err := ctx.Res.Write(value)
	return err
}
func (ctx *Context) TEXT(value string) error {
	ctx.SetHeader("Content-Type", "text/plain; charset=utf-8")
	return ctx.BODY([]byte(value))
}

func (ctx *Context) JSON(value any) error {
	marshal, err := json.Marshal(value)
	if err != nil {
		return err
	}
	ctx.SetHeader("Content-Type", "application/json; charset=utf-8")
	return ctx.BODY(marshal)
}

func (ctx *Context) HTML(value string) error {
	ctx.SetHeader("Content-Type", "text/html; charset=utf-8")
	return ctx.BODY([]byte(value))
}

