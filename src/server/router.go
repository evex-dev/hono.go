package server

import (
	"strings"
	"sync"
)

var wg sync.WaitGroup

type FoundRoutes []*Route

func NewRouter(routes *Routes) *Router {
	trie := NewTrieTree(routes).Build()

	return &Router{
		Match: func (pattern string, method string) FoundRoutes {
			foundRoutes := FoundRoutes{}
			isAleadyFound := false

			for _, grafting := range trie.Tree {
				if isPassable(grafting, pattern) && (method == "ALL" || method == grafting.Route.Method) {
					if isAleadyFound && !grafting.Route.IsMiddleware {
						continue
					}

					foundRoutes = append(foundRoutes, grafting.Route)

					if !isAleadyFound && !grafting.Route.IsMiddleware {
						isAleadyFound = true
					}
				}
			}

			if len(foundRoutes) != 0 {
				return foundRoutes
			}

			return nil
		},
	}
}

type TrieGrafting struct {
	Route    *Route
	TreePath []string
}

type TrieTree struct {
	Initlized bool
	Tree      []*TrieGrafting
	Routes
}

func NewTrieTree(routes *Routes) *TrieTree {
	return &TrieTree{
		Routes:    *routes,
		Initlized: false,
	}
}

func (t *TrieTree) Build() *TrieTree {

	for _, route := range t.Routes.RouteList {
		t.Tree = append(t.Tree, &TrieGrafting{
			Route:    route,
			TreePath: strings.Split(PathFixer(route.Pattern), "/")[1:],
		})
	}

	t.Initlized = true
	return t
}

func PathFixer(path string) string {
	path = strings.Trim(path, " ")

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return path
}

const ALL_METHODS = "ALL"

func isPassable(trieGrafting *TrieGrafting, pattern string) bool {
	treePath := strings.Split(PathFixer(pattern), "/")[1:]

	
}
