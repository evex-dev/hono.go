package context

func (ctx *Context) SetHeader(key, value string) *Context {
	ctx.Res.Header().Set(key, value)
	return ctx
}

func (ctx *Context) AddHeader(key, value string) *Context {
	ctx.Res.Header().Add(key, value)
	return ctx
}

func (ctx *Context) Status(status int) *Context {
	ctx.Res.WriteHeader(status)
	return ctx
}
