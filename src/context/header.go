package context

func (ctx *Context) SetHeader(key, value string) {
	ctx.Res.Header().Set(key, value)
}

func (ctx *Context) AddHeader(key, value string) {
	ctx.Res.Header().Add(key, value)
}

func (ctx *Context) Status(status int) {
	ctx.Res.WriteHeader(status)
}
