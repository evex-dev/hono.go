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
	Port        string
	Host        string
	isPortReady bool
	isHostReady bool
	Err         error
	Fire        func() error
}

func (r *RunContext) Callback(callbackFunc func(addr string, err error) error) *RunContext {
	callbackResult := callbackFunc(r.Host+r.Port, r.Err)
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
	c := &RunContext{
		Port:        ":3000",
		Host:        "0.0.0.0",
		isPortReady: true,
		isHostReady: false,
	}

	c.Fire = func() error {
		return h.Engine.Run(c.Host + c.Port)
	}

	return c
}

func (h *HonoGo) InitTLS(cert, key string) *RunContext {
	c := &RunContext{
		Host: "0.0.0.0",
		Port: ":3000",
	}

	c.Fire = func() error {
		return h.Engine.RunTLS(c.Host+c.Port, cert, key)
	}

	return c
}

// NotFound And Middleware
func (h *HonoGo) NotFound(handler HandlerFunc) *HonoGo {
	h.Engine.NotFoundHandler = handler
	return h
}

func (h *HonoGo) Use(pattern string, handlers ...HandlerFunc) *HonoGo {
	for i := 0; i < len(handlers); i++ {
		h.Engine.AddRoute("ALL", pattern, handlers[i], true)
	}

	return h
}

// Methods

func (h *HonoGo) Get(pattern string, handlers ...HandlerFunc) *HonoGo {
	for i := 0; i < len(handlers); i++ {
		h.Engine.AddRoute("GET", pattern, handlers[i], false)
	}
	return h
}

func (h *HonoGo) Head(pattern string, handlers ...HandlerFunc) *HonoGo {
	for i := 0; i < len(handlers); i++ {
		h.Engine.AddRoute("GET", pattern, handlers[i], false)
	}
	return h
}

func (h *HonoGo) Post(pattern string, handlers ...HandlerFunc) *HonoGo {
	for i := 0; i < len(handlers); i++ {
		h.Engine.AddRoute("POST", pattern, handlers[i], false)
	}
	return h
}

func (h *HonoGo) Put(pattern string, handlers ...HandlerFunc) *HonoGo {
	for i := 0; i < len(handlers); i++ {
		h.Engine.AddRoute("PUT", pattern, handlers[i], false)
	}
	return h
}

func (h *HonoGo) Delete(pattern string, handlers ...HandlerFunc) *HonoGo {
	for i := 0; i < len(handlers); i++ {
		h.Engine.AddRoute("DELETE", pattern, handlers[i], false)
	}
	return h
}

func (h *HonoGo) Options(pattern string, handlers ...HandlerFunc) *HonoGo {
	for i := 0; i < len(handlers); i++ {
		h.Engine.AddRoute("OPTIONS", pattern, handlers[i], false)
	}
	return h
}

func (h *HonoGo) All(pattern string, handlers ...HandlerFunc) *HonoGo {
	for i := 0; i < len(handlers); i++ {
		h.Engine.AddRoute("ALL", pattern, handlers[i], false)
	}
	return h
}

func (h *HonoGo) On(method string, pattern string, handlers ...HandlerFunc) *HonoGo {
	for i := 0; i < len(handlers); i++ {
		h.Engine.AddRoute(method, pattern, handlers[i], false)
	}
	return h
}
