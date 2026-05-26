package gen

import (
	"go.uber.org/zap"

	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/openapi"
)

// Generator is OpenAPI-to-Go generator.
type Generator struct {
	opt               GenerateOptions
	parseOpts         ParseOptions
	api               *openapi.API
	servers           []ir.Server
	operations        []*ir.Operation
	defaultOperations []*ir.Operation // Operations without an operation group.
	operationGroups   []*ir.OperationGroup
	webhooks          []*ir.Operation
	securities        map[string]*ir.Security
	tstorage          *tstorage
	errType           *ir.Response
	webhookRouter     WebhookRouter
	router            Router
	imports           map[string]string
	equalitySpecs     []*ir.EqualityMethodSpec // Types requiring Equal() methods for uniqueItems validation

	log *zap.Logger
}

func expandSpec(api *openapi.API, p string) (err error) { _ = "STUB: not implemented"; return nil }

// NewGenerator creates new Generator.
func NewGenerator(spec *ogen.Spec, opts Options) (*Generator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Default: allow cross-type constraints unless explicitly set to false

func (g *Generator) makeIR(api *openapi.API) error { _ = "STUB: not implemented"; return nil }

// Collect types that need Equal() and Hash() methods for complex uniqueItems validation

func (g *Generator) makeServers(servers []openapi.Server) error {
	_ = "STUB: not implemented"
	return nil

	// Ignore servers without name.
}

func (g *Generator) makeOps(ops []*openapi.Operation) error { _ = "STUB: not implemented"; return nil }

func (g *Generator) makeWebhooks(webhooks []openapi.Webhook) error {
	_ = "STUB: not implemented"
	return nil
}

func sortOperations(ops []*ir.Operation) { _ = "STUB: not implemented"; return }

func groupOperations(ops []*ir.Operation) (
	defaultOperations []*ir.Operation,
	operationGroups []*ir.OperationGroup,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Types returns generated types.
func (g *Generator) Types() map[string]*ir.Type { _ = "STUB: not implemented"; return nil }

// Operations returns generated operations.
func (g *Generator) Operations() []*ir.Operation { _ = "STUB: not implemented"; return nil }

// Webhooks returns generated webhooks.
func (g *Generator) Webhooks() []*ir.Operation {
	_ = "STUB: not implemented"

	// API returns api schema.
	return nil
}

func (g *Generator) API() *openapi.API { _ = "STUB: not implemented"; return nil }
