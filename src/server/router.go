package server

func NewRouter(routes *Routes) *Router {
	return &Router{
		Match: func(path string) *Route {
			for _, route := range routes.Route {
				if route.Pattern == path {
					return route
				}
			}
			return nil
		},
	}
}
