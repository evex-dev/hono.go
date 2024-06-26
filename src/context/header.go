package context

func (ctx *Context) SetHeader(key, value string) {
	ctx.Writer.Header().Set(key, value)
}

func (ctx *Context) AddHeader(key, value string) {
	ctx.Writer.Header().Add(key, value)
}

func (ctx *Context) Status(status int) {
	ctx.Writer.WriteHeader(status)
}
