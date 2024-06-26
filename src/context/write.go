package context

func (ctx *Context) Write(value []byte) error {
	_, err := ctx.Writer.Write(value)
	return err
}
