package gen

import (
	"github.com/ogen-go/ogen/gen/ir"
)

// RouteTree is Radix tree of routes.
type RouteTree struct {
	Root *RouteNode
}

func findParam(op *ir.Operation, name string) (*ir.Parameter, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// longestPrefix founds the longest common prefix of k1 and k2.
func longestPrefix(k1, k2 string) int { _ = "STUB: not implemented"; return 0 }

func (t *RouteTree) addRoute(m Route) error { _ = "STUB: not implemented"; return nil }

// Head is a first character of route.

// Find parameter index.

// Check for existing node with same head.

// If there is no child with such head, create a new one.

// Skip common parameter node.
//
// This condition is met if child have parameter part in same place, e.g.
//
// /pet/{name}
// /pet/{name}/friends
//

// Found the longest common prefix of existing node and new.

// If common prefix fully matched by existing node, try to create child.

// Otherwise, we try to replace existing node.

// Add existing node as child of replacer.

// Special case: if new node has exactly same path, replace existing node.

// Otherwise, we add new node as second child of replacer.

func (t *RouteTree) Walk(cb func(level int, n *RouteNode)) { _ = "STUB: not implemented"; return }
