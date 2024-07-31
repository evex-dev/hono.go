package server

type HonoGo struct {
	Engine *Engine
}

func Create() *HonoGo {
	return &HonoGo{
		Engine: &Engine{},
	}
}

// Listen function

type RunContext struct {
	Port string
	Host string
	isPortReady bool
	isHostReady bool
	Err error
	Fire func() error
}

func (r *RunContext) Callback(callbackFunc func(addr string, err error) error) *RunContext {
	callbackResult := callbackFunc(r.Host + r.Port, r.Err)
	if callbackResult != nil {
		r.Err = callbackResult
	}
	return r
}

func (r *RunContext) SetPort(port string) *RunContext {
	r.Port = ":" + port
	return r
}

func (r *RunContext) SetHost(host string) *RunContext {
	r.Host = host
	return r
}

// Init

func (h *HonoGo) Init() *RunContext {
	ctx := &RunContext{
		Port: ":3000",
		Host: "0.0.0.0",
		isPortReady: true,
		isHostReady: false,
	}

	ctx.Fire = func() error {
		return h.Engine.Run(ctx.Host + ctx.Port)
	}

	return ctx
}

func (h *HonoGo) InitTLS(cert, key string) *RunContext {
	ctx := &RunContext{
		Host: "0.0.0.0",
		Port: ":3000",
	}

	ctx.Fire = func() error {
		return h.Engine.RunTLS(ctx.Host + ctx.Port, cert, key)
	}

	return ctx
}

// NotFound And Middleware
func (h *HonoGo) NotFound(handler HandlerFunc) *HonoGo {
	h.Engine.NotFoundHandler = handler
	return h
}

func (h *HonoGo) Use(pattern string, handler HandlerFunc) *HonoGo {
	h.Engine.AddRoute("ALL", pattern, handler, true)
	return h
}

// Methods

func (h *HonoGo) Get(pattern string, handler HandlerFunc) *HonoGo {
	h.Engine.AddRoute("GET", pattern, handler, false)
	return h
}

func (h *HonoGo) Head(pattern string, handler HandlerFunc) *HonoGo {
	h.Engine.AddRoute("HEAD", pattern, handler, false)
	return h
}

func (h *HonoGo) Post(pattern string, handler HandlerFunc) *HonoGo {
	h.Engine.AddRoute("POST", pattern, handler, false)
	return h
}

func (h *HonoGo) Put(pattern string, handler HandlerFunc) *HonoGo {
	h.Engine.AddRoute("PUT", pattern, handler, false)
	return h
}

func (h *HonoGo) Delete(pattern string, handler HandlerFunc) *HonoGo {
	h.Engine.AddRoute("DELETE", pattern, handler, false)
	return h
}

func (h *HonoGo) Options(pattern string, handler HandlerFunc) *HonoGo {
	h.Engine.AddRoute("OPTIONS", pattern, handler, false)
	return h
}

func (h *HonoGo) All(pattern string, handler HandlerFunc) *HonoGo {
	h.Engine.AddRoute("ALL", pattern, handler, false)
	return h
}

func (h *HonoGo) On(method string, pattern string, handler HandlerFunc) *HonoGo {
	h.Engine.AddRoute(method, pattern, handler, false)
	return h
}
