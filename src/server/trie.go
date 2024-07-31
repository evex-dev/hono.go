package server

import (
	"regexp"
	"strings"

	"github.com/evex-dev/hono.go/src/context"
)

type TrieFunc func(params map[string]string) (string, map[string]string)

type Pattern struct {
	Key     string
	Name    string
	Matcher *regexp.Regexp
}

type TrieNode struct {
	Methods  map[string]TrieFunc
	Children map[string]*TrieNode
	Patterns []Pattern
	Order    int
	Name     string
	Params   map[string]string
}

type ParamsValues struct {
	Values []ParamsValue
}

type ParamsValue struct {
	Path   string
	Params map[string]string
}

type Value struct {
	Handler TrieFunc
	Params  map[string]string
	Name    string
	Score   int
}

func NewTrie() *TrieNode {
	return &TrieNode{
		Children: make(map[string]*TrieNode),
		Methods:  make(map[string]TrieFunc),
		Params:   make(map[string]string),
	}
}

func (t *TrieNode) getHandlerSets(nodeParams, params map[string]string) []Value {
	var handlerSets []Value

	for _, handler := range t.Methods {
		handlerSet := Value{
			Handler: handler,
			Params:  make(map[string]string),
			Name:    t.Name,
			Score:   t.Order,
		}

		for key, value := range nodeParams {
			handlerSet.Params[key] = value
		}
		for key, value := range params {
			handlerSet.Params[key] = value
		}

		handlerSets = append(handlerSets, handlerSet)
	}

	return handlerSets
}

// AddRoute => Insert
func (t *TrieNode) Insert(path string) {
	t.Name = path
	parts := splitRoutingPath(path)
	curNode := t

	for i, p := range parts {
		if p == "*" {
			if _, exists := curNode.Children["*"]; !exists {
				curNode.Children["*"] = NewTrie()
			}
			curNode = curNode.Children["*"]
			if i == len(parts)-1 {
				continue
			}
			continue
		}

		if child, exists := curNode.Children[p]; exists {
			curNode = child
			continue
		}

		child := NewTrie()
		if pattern := getPattern(p); pattern != nil {
			curNode.Patterns = append(curNode.Patterns, *pattern)
		}
		curNode.Children[p] = child
		curNode = child
	}

	curNode.Methods["params"] = func(params map[string]string) (string, map[string]string) {
		return path, params
	}
}

func (t *TrieNode) Search(path string) []Value {
	var sets []Value
	t.Params = make(map[string]string)

	curNodes := []*TrieNode{t}
	parts := splitPath(path)

	for i, part := range parts {
		isLast := i == len(parts)-1
		tempNodes := []*TrieNode{}

		for _, node := range curNodes {
			if nextNode, exists := node.Children[part]; exists {
				nextNode.Params = node.Params

				if isLast {
					sets = append(sets, nextNode.getHandlerSets(nextNode.Params, map[string]string{})...)
					if childNode, exists := nextNode.Children["*"]; exists {
						sets = append(sets, childNode.getHandlerSets(nextNode.Params, map[string]string{})...)
					}
				} else {
					tempNodes = append(tempNodes, nextNode)
				}
			}

			for _, pattern := range node.Patterns {
				if pattern.Matcher != nil {
					if pattern.Matcher.MatchString(part) {
						params := make(map[string]string)
						params[pattern.Name] = part

						if isLast {
							sets = append(sets, node.getHandlerSets(params, node.Params)...)
							if childNode, exists := node.Children[pattern.Key]; exists {
								sets = append(sets, childNode.getHandlerSets(params, node.Params)...)
							}
						} else {
							if childNode, exists := node.Children[pattern.Key]; exists {
								childNode.Params = params
							}
						}
					}
				} else {
					params := make(map[string]string)
					params[pattern.Name] = part

					if isLast {
						sets = append(sets, node.getHandlerSets(params, node.Params)...)
						if childNode, exists := node.Children[pattern.Key]; exists {
							sets = append(sets, childNode.getHandlerSets(params, node.Params)...)
						}
					} else {
						if childNode, exists := node.Children[pattern.Key]; exists {
							childNode.Params = params
						}
					}
				}
			}

			if wildcardNode, exists := node.Children["*"]; exists {
				params := make(map[string]string)
				if isLast {
					sets = append(sets, wildcardNode.getHandlerSets(params, node.Params)...)
				} else {
					tempNodes = append(tempNodes, wildcardNode)
				}
			}
		}

		curNodes = tempNodes
	}

	return sortHandlerSets(sets)
}

func sortHandlerSets(sets []Value) []Value {
	for i := 0; i < len(sets)-1; i++ {
		for j := i + 1; j < len(sets); j++ {
			if sets[i].Score > sets[j].Score {
				sets[i], sets[j] = sets[j], sets[i]
			}
		}
	}
	return sets
}

func splitPath(path string) []string {
	return strings.Split(strings.Trim(path, "/"), "/")
}

func splitRoutingPath(path string) []string {
	return splitPath(path)
}

func getPattern(part string) *Pattern {
	if strings.HasPrefix(part, ":") {
		paramName := strings.TrimPrefix(part, ":")
		var matcher *regexp.Regexp
		if strings.Contains(paramName, "{") && strings.Contains(paramName, "}") {
			parts := strings.Split(paramName, "{")
			regexPart := strings.TrimSuffix(parts[1], "}")
			matcher = regexp.MustCompile(regexPart)
			paramName = parts[0]
		}
		return &Pattern{
			Key:     part,
			Name:    paramName,
			Matcher: matcher,
		}
	}
	return nil
}

// GetRoute => Find
func (t *TrieNode) Find(path string) *ParamsValues {
	sets := t.Search(path)
	var values ParamsValues
	for _, set := range sets {
		path, pattern := set.Handler(set.Params)
		values.Values = append(values.Values, ParamsValue{
			Path:   path,
			Params: pattern,
		})
	}
	return &values
}

// Router
// experimental support for /* by evorax
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

func (t *TrieNode) Compare(path, dst string) (*context.Params, bool) {
	values := t.Find(dst)

	params := make(context.Params)

	for _, value := range values.Values {
		for key, value2 := range value.Params {
			params[key] = value2
		}

		return &params, value.Path == path
	}

	return nil, false
}
