package parser

import (
	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/jsonpointer"
	"github.com/ogen-go/ogen/openapi"
)

func (p *parser) parseWebhook(name string, item *ogen.PathItem, ctx *jsonpointer.ResolveCtx) (openapi.Webhook, error) {
	_ = "STUB: not implemented"
	// FIXME(tdakkota): we are passing "/" path, but webhook has no path.
	return *new(openapi.Webhook), nil
}

func (p *parser) parseWebhooks(webhooks map[string]*ogen.PathItem) (r []openapi.Webhook, rerr error) {
	_ = "STUB: not implemented"
	return nil, nil
}
