package server

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/evex-dev/hono.go/src/context"
)

// structs
type TrieNode struct {
	children map[string]*TrieNode
	isEnd    bool
	paramKey string
	regex    *regexp.Regexp
	wildcard *TrieNode
}

type Trie struct {
	root     *TrieNode
	patterns map[*TrieNode]string
}

// trie tree
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[string]*TrieNode),
		},
		patterns: make(map[*TrieNode]string),
	}
}

func (t *Trie) AddRoute(pattern string) {
	parts := strings.Split(pattern, "/")[1:]
	node := t.root

	for _, part := range parts {
		if part == "*" {
			if node.wildcard == nil {
				node.wildcard = &TrieNode{children: make(map[string]*TrieNode)}
			}
			node = node.wildcard
		} else if len(part) > 0 && part[0] == ':' {
			paramKey := part[1:]
			if openBracketIndex := strings.Index(paramKey, "{"); openBracketIndex != -1 {
				regexPattern := paramKey[openBracketIndex+1 : len(paramKey)-1]
				paramKey = paramKey[:openBracketIndex]
				node = t.AddChild(node, paramKey, regexPattern)
			} else {
				node = t.AddChild(node, paramKey, "")
			}
		} else {
			if _, ok := node.children[part]; !ok {
				node.children[part] = &TrieNode{children: make(map[string]*TrieNode)}
			}
			node = node.children[part]
		}
	}

	node.isEnd = true
	t.patterns[node] = pattern
}

func (t *Trie) AddChild(node *TrieNode, paramKey string, regexPattern string) *TrieNode {
	for _, child := range node.children {
		if child.paramKey == paramKey {
			return child
		}
	}

	newNode := &TrieNode{
		children: make(map[string]*TrieNode),
		paramKey: paramKey,
	}

	if regexPattern != "" {
		newNode.regex = regexp.MustCompile(regexPattern)
	}

	node.children[paramKey] = newNode
	return newNode
}

func (t *Trie) GetRoute(path string) (string, *context.Params) {
	parts := strings.Split(path, "/")[1:]
	node := t.root

	params := make(context.Params)
	var matchedPattern string

	for _, part := range parts {
		if _, ok := node.children[part]; ok {
			node = node.children[part]
		} else if node.wildcard != nil {
			node = node.wildcard
		} else {
			matched := false
			for _, child := range node.children {
				if child.paramKey != "" {
					if child.regex != nil && child.regex.MatchString(part) {
						params[child.paramKey] = part
						node = child
						matched = true
						break
					} else if child.regex == nil {
						params[child.paramKey] = part
						node = child
						matched = true
						break
					}
				}
			}
			if !matched {
				return "", nil
			}
		}
	}

	if node.isEnd {
		matchedPattern = t.patterns[node]
	} else if node.wildcard != nil {
		matchedPattern = "/*"
	}

	return matchedPattern, &params
}

// Router
// /* is not support, so coming soon by evorax
func NewTrieRouter(routes *Routes) *Router {
	trie := NewTrie()

	for _, route := range routes.RouteList {
		trie.AddRoute(route.Pattern)
	}

	return &Router{
		Match: func(pattern string, method string) (FoundRoutes, *context.Params) {
			foundRoutes := FoundRoutes{}
			params := &context.Params{}
			isFound := false
			fixedPattern := PathFixer(pattern)

			for _, route := range routes.RouteList {
				if route.Method != ALL_METHODS && route.Method != method {
					continue
				}

				paramsResult, isSuccess := trie.Compare(PathFixer(route.Pattern), fixedPattern)
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

func (t *Trie) Compare(path, dst string) (*context.Params, bool) {
	pattern, params := t.GetRoute(dst)

	// debug
	fmt.Println("[DEBUG]", pattern, *params, pattern == path, fmt.Sprintf("(%s == %s)", pattern, path))

	return params, pattern == path
}
