package context

import (
	"encoding/json"
	// "encoding/xml"
	"fmt"
)

func (ctx *Context) Write(value []byte) error {
	_, err := ctx.Writer.Write(value)
	return err
}
func (ctx *Context) WriteString(value string) error {
	return ctx.Write([]byte(value))
}

func (ctx *Context) Fprint(value string) error {
	_, err := fmt.Fprint(ctx.Writer, value)
	return err
}

func (ctx *Context) Fprintf(value string) error {
	_, err := fmt.Fprintf(ctx.Writer, value)
	return err
}

func (ctx *Context) Fprintln(value string) error {
	_, err := fmt.Fprintln(ctx.Writer, value)
	return err
}

func (ctx *Context) JSON(value any) error {
	marshal, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return ctx.Write(marshal)
}
