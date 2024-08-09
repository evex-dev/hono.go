package context

func (c *Context) SetHeader(key, value string) *Context {
	c.Res.Header().Set(key, value)
	return c
}

func (c *Context) AddHeader(key, value string) *Context {
	c.Res.Header().Add(key, value)
	return c
}

func (c *Context) Status(status int) *Context {
	c.Res.WriteHeader(status)
	return c
}
