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

func (h *HonoGo) Run(addr string) error {
	return h.Engine.Run(addr)
}

// func (h *HonoGo) RunTLS(addr string, cert, key string) error {
// 	return h.Engine.Run(addr)
// }

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
