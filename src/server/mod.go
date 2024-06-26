package server

type HonoGo struct {
	Engine *Engine
}

func NewHonoGo() *HonoGo {
	return &HonoGo{
		Engine: &Engine{},
	}
}

// Listen function

type RunContext struct {
	Addr string
	Err error
	Fire func() error
}

func (r *RunContext) Callback(callbackFunc func(addr string, err error) error) *RunContext {
	callbackResult := callbackFunc(r.Addr, r.Err)
	if callbackResult != nil {
		r.Err = callbackResult
	}
	return r
}

func (r *RunContext) SetPort(addr string) *RunContext {
	r.Addr = ":" + addr
	return r
}

// Init

func (h *HonoGo) Init() *RunContext {
	ctx := &RunContext{
		Addr: ":3000",
	}

	ctx.Fire = func() error {
		return h.Engine.Run(ctx.Addr)
	}

	return ctx
}

func (h *HonoGo) InitTLS(cert, key string) *RunContext {
	ctx := &RunContext{
		Addr: ":3000",
	}

	ctx.Fire = func() error {
		return h.Engine.RunTLS(ctx.Addr, cert, key)
	}

	return ctx
}

// Methods

func (h *HonoGo) GET(pattern string, handler HandlerFunc) *HonoGo {
	h.Engine.AddRoute("GET", pattern, handler)
	return h
}

func (h *HonoGo) HEAD(pattern string, handler HandlerFunc) *HonoGo {
	h.Engine.AddRoute("HEAD", pattern, handler)
	return h
}

func (h *HonoGo) POST(pattern string, handler HandlerFunc) *HonoGo {
	h.Engine.AddRoute("POST", pattern, handler)
	return h
}

func (h *HonoGo) PUT(pattern string, handler HandlerFunc) *HonoGo {
	h.Engine.AddRoute("PUT", pattern, handler)
	return h
}

func (h *HonoGo) DELETE(pattern string, handler HandlerFunc) *HonoGo {
	h.Engine.AddRoute("DELETE", pattern, handler)
	return h
}

func (h *HonoGo) OPTIONS(pattern string, handler HandlerFunc) *HonoGo {
	h.Engine.AddRoute("OPTIONS", pattern, handler)
	return h
}

func (h *HonoGo) ON(method string, pattern string, handler HandlerFunc) *HonoGo {
	h.Engine.AddRoute(method, pattern, handler)
	return h
}
