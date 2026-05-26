package gen

import (
	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/openapi"
)

// Route describes route.
type Route struct {
	Method    string        // GET, POST, DELETE
	Path      string        // /api/v1/user/{name}/info
	Operation *ir.Operation // getUserInfo
}

// Routes is list of routes.
type Routes []Route

// Len implements sort.Interface.
func (n Routes) Len() int {
	_ = "STUB: not implemented"

	// Less implements sort.Interface.
	return 0
}

func (n Routes) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap implements sort.Interface.
func (n Routes) Swap(i, j int) { _ = "STUB: not implemented"; return }

// AddRoute adds new route. If the route is already added, it returns error.
func (n *Routes) AddRoute(nr Route) error { _ = "STUB: not implemented"; return nil }

// Keep routes sorted by method.

// Router contains list of routes.
type Router struct {
	Tree RouteTree
	// MaxParametersCount is maximum number of path parameters in one operation.
	MaxParametersCount int
}

// Add adds new route.
func (s *Router) Add(r Route) error { _ = "STUB: not implemented"; return nil }

// NodesWithAllowedHeaders returns route nodes that
// contain routes that accept headers.
func (s Router) NodesWithAllowedHeaders() []*RouteNode { _ = "STUB: not implemented"; return nil }

// WebhookRoute is a webhook route.
type WebhookRoute struct {
	Method    string
	Operation *ir.Operation
}

// WebhookRoutes is a list of webhook methods.
type WebhookRoutes struct {
	Routes []WebhookRoute
	idSeq  idSeq
}

// ID returns list identifier.
func (r WebhookRoutes) ID() int {
	_ = "STUB: not implemented"

	// Add adds new operation to the route.
	return 0
}

func (r *WebhookRoutes) Add(nr WebhookRoute) error { _ = "STUB: not implemented"; return nil }

// AllowedMethods returns comma-separated list of allowed methods.
func (r WebhookRoutes) AllowedMethods() string { _ = "STUB: not implemented"; return "" }

// WithAllowedHeaders reports whether any route
// in this list accepts any headers.
func (r WebhookRoutes) WithAllowedHeaders() bool { _ = "STUB: not implemented"; return false }

// AllowedHeaders returns HTTP method and allowed headers pairs.
// Allowed headers are formatted as a comma-separated list of headers.
func (r WebhookRoutes) AllowedHeaders() [][2]string { _ = "STUB: not implemented"; return nil }

// PostContentTypes returns comma-separated list of content types
// accepted by a route with POST method.
func (r WebhookRoutes) PostContentTypes() string { _ = "STUB: not implemented"; return "" }

// PatchContentTypes returns comma-separated list of content types
// accepted by a route with PATCH method.
func (r WebhookRoutes) PatchContentTypes() string { _ = "STUB: not implemented"; return "" }

// methodContentTypes returns comma-separated list of content types
// accepted by a route with the specified HTTP method.
func (r WebhookRoutes) methodContentTypes(method string) string {
	_ = "STUB: not implemented"
	return ""
}

// WebhookRouter contains routing information for webhooks.
type WebhookRouter struct {
	Webhooks map[string]WebhookRoutes
	idSeq    idSeq
}

// WebhooksWithAllowedHeaders returns lists that
// contain routes that accept any headers.
func (r WebhookRouter) WebhooksWithAllowedHeaders() []WebhookRoutes {
	_ = "STUB: not implemented"
	return nil
}

// Add adds new route.
func (r *WebhookRouter) Add(name string, nr WebhookRoute) error {
	_ = "STUB: not implemented"
	return nil
}

func checkRoutePath(p openapi.Path) error { _ = "STUB: not implemented"; return nil }

// Cond: i > 0

func (g *Generator) route() error { _ = "STUB: not implemented"; return nil }
