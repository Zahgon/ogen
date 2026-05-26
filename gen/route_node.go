package gen

import (
	"github.com/ogen-go/ogen/gen/ir"
)

// idSeq is a monotonically increasing id sequence.
type idSeq struct {
	id  int
	seq *int
}

func (s *idSeq) next() idSeq { _ = "STUB: not implemented"; return *new(idSeq) }

type nodes []*RouteNode

func (e nodes) Len() int { _ = "STUB: not implemented"; return 0 }

func (e nodes) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (e nodes) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (e nodes) Sort() {
	_ = "STUB: not implemented"

	// RouteNode is node of Radix tree of routes.
	return
}

type RouteNode struct {
	idSeq idSeq

	prefix string
	head   byte
	child  nodes

	paramName string
	param     *ir.Parameter // nil-able

	routes Routes
}

// ID returns node identifier.
func (n *RouteNode) ID() int {
	_ = "STUB: not implemented"

	// AddRoute adds new method route to node.
	return 0
}

func (n *RouteNode) AddRoute(nr Route) error { _ = "STUB: not implemented"; return nil }

// Prefix returns common prefix.
func (n *RouteNode) Prefix() string {
	_ = "STUB: not implemented"

	// Head returns first byte of prefix.
	return ""
}

func (n *RouteNode) Head() byte {
	_ = "STUB: not implemented"

	// IsStatic whether node is not a parameter node.
	return 0
}

func (n *RouteNode) IsStatic() bool { _ = "STUB: not implemented"; return false }

// IsLeaf whether node has no children.
func (n *RouteNode) IsLeaf() bool { _ = "STUB: not implemented"; return false }

// IsParam whether node is a parameter node.
func (n *RouteNode) IsParam() bool { _ = "STUB: not implemented"; return false }

// Children returns child nodes.
func (n *RouteNode) Children() []*RouteNode {
	_ = "STUB: not implemented"

	// StaticChildren returns slice of child static nodes.
	return nil
}

func (n *RouteNode) StaticChildren() (r []*RouteNode) { _ = "STUB: not implemented"; return nil }

// ParamChildren returns slice of child parameter nodes.
func (n *RouteNode) ParamChildren() (r []*RouteNode) { _ = "STUB: not implemented"; return nil }

// Tails returns heads of child nodes.
//
// Used for matching end of parameter node between two static.
func (n *RouteNode) Tails() (r []byte) { _ = "STUB: not implemented"; return nil }

// ParamName returns parameter name, if any.
func (n *RouteNode) ParamName() string {
	_ = "STUB: not implemented"

	// Param returns associated parameter, if any.
	//
	// May be nil.
	return ""
}

func (n *RouteNode) Param() *ir.Parameter {
	_ = "STUB: not implemented"

	// AllowedMethods returns list of allowed methods.
	return nil
}

func (n *RouteNode) AllowedMethods() string { _ = "STUB: not implemented"; return "" }

// WithAllowedHeaders reports whether any route
// in this node accepts any headers.
func (n *RouteNode) WithAllowedHeaders() bool { _ = "STUB: not implemented"; return false }

// AllowedHeaders returns HTTP method and allowed headers pairs.
// Allowed headers are formatted as a comma-separated list of headers.
func (n *RouteNode) AllowedHeaders() [][2]string { _ = "STUB: not implemented"; return nil }

// PostContentTypes returns comma-separated list of content types
// accepted by a route with POST method.
func (n *RouteNode) PostContentTypes() string { _ = "STUB: not implemented"; return "" }

// PatchContentTypes returns comma-separated list of content types
// accepted by a route with PATCH method.
func (n *RouteNode) PatchContentTypes() string { _ = "STUB: not implemented"; return "" }

// methodContentTypes returns comma-separated list of content types
// accepted by a route with the specified HTTP method.
func (n *RouteNode) methodContentTypes(method string) string { _ = "STUB: not implemented"; return "" }

// Routes returns list of associated Route.
func (n *RouteNode) Routes() []Route { _ = "STUB: not implemented"; return nil }

func nextPathPart(s string) (hasParam bool, paramStart, paramEnd int, _ error) {
	_ = "STUB: not implemented"
	return false, 0, 0, nil
}

// Need to match parameter part including both brackets.

func (n *RouteNode) addChild(path string, op *ir.Operation, ch *RouteNode) (r *RouteNode, _ error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Route starts with a param.

// Handle tail of path.

// Route contains param.
// Set prefix to static part of path.

// Get parameterized part.

// Add parameterized child node.

func (n *RouteNode) childIdx(head byte) (int, bool) { _ = "STUB: not implemented"; return 0, false }

func (n *RouteNode) replaceChild(head byte, child *RouteNode) { _ = "STUB: not implemented"; return }

func (n *RouteNode) getChild(head byte) *RouteNode { _ = "STUB: not implemented"; return nil }

func (n *RouteNode) walk(level int, cb func(level int, n *RouteNode)) {
	_ = "STUB: not implemented"
	return
}
