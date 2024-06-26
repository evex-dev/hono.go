package context

func (ctx *Context) Write(value []byte) error {
	_, err := ctx.Writer.Write(value)
	return err
}
func (ctx *Context) WriteString(value []byte) error {
	return Wr
}
