package server

import (
	"github.com/evex-dev/hono.go/src/context"
	"regexp"
	"strings"
)

type TrieNode struct {
	children      map[string]*TrieNode
	isEnd         bool
	paramChild    *TrieNode
	regexpChild   *TrieNode
	regexpPattern *regexp.Regexp
	paramName     string
	pattern       string
}

type Trie struct {
	root *TrieNode
}

func NewTrieRouter(routes *Routes) *Router {

	trie := NewTrie()

	for _, route := range routes.RouteList {
		trie.Insert(route.Pattern)
	}

	return &Router{
		Match: func(pattern string, method string) (FoundRoutes, *context.Params) {
			foundRoutes := FoundRoutes{}
			params := &context.Params{}
			isFound := false
			
			for _, route := range routes.RouteList {
				if route.Method != method && method != "ALL" {
					continue
				}

				paramsResult, isSuccess := trie.Compare(route.Pattern, pattern)
				if isSuccess {
					params = MergeParams(params, paramsResult)
					foundRoutes = append(foundRoutes, route)
					isFound = true
				}
			}

			if isFound {
				return foundRoutes, params
			}

			return nil, params
		},
	}
}

func MergeParams(params1, params2 *context.Params) *context.Params {
	for key, value := range *params2 {
		(*params1)[key] = value
	}
	return params1
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[string]*TrieNode),
		},
	}
}

func (t *Trie) Insert(path string) {
	current := t.root

	parts := strings.Split(path, "/")
	for _, part := range parts {
		if part == "" {
			continue
		}

		if strings.HasPrefix(part, ":") && strings.Contains(part, "{") && strings.Contains(part, "}") {
			paramName := part[1:strings.Index(part, "{")]
			regexpPattern := part[strings.Index(part, "{")+1 : strings.Index(part, "}")]
			if current.paramChild == nil {
				current.paramChild = &TrieNode{
					children:  make(map[string]*TrieNode),
					paramName: paramName,
				}
			}
			current = current.paramChild
			current.regexpPattern = regexp.MustCompile(regexpPattern)
		} else if strings.HasPrefix(part, ":") {
			paramName := part[1:]
			if current.paramChild == nil {
				current.paramChild = &TrieNode{
					children:  make(map[string]*TrieNode),
					paramName: paramName,
				}
			}
			current = current.paramChild
			current.paramName = paramName
		} else {
			if current.children[part] == nil {
				current.children[part] = &TrieNode{
					children: make(map[string]*TrieNode),
				}
			}
			current = current.children[part]
		}
	}
	current.isEnd = true
	current.pattern = path
}

func (t *Trie) Compare(path, pattern string) (*context.Params, bool) {
	param, match, ok := t.Match(path)
	if !ok {
		return nil, ok
	}

	if match == pattern {
		return nil, false
	}

	return param, true
}

func (t *Trie) Match(path string) (*context.Params, string, bool) {
	current := t.root

	params := make(context.Params)

	var match string

	parts := strings.Split(path, "/")
	for _, part := range parts {
		if part == "" {
			continue
		}

		if current.children[part] != nil {
			current = current.children[part]
		} else if current.paramChild != nil {
			current = current.paramChild
			params[current.paramName] = part
		} else if current.regexpChild != nil {
			current = current.regexpChild
			if current.regexpPattern.MatchString(part) {
				params[current.paramName] = part
			} else {
				return nil, "", false
			}
		} else {
			return nil, "", false
		}
	}

	if current.isEnd {
		match = current.pattern
	}

	return &params, match, current.isEnd
}
