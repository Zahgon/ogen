package parser

import (
	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/jsonpointer"
	"github.com/ogen-go/ogen/openapi"
)

func (p *parser) parseHeaders(headers map[string]*ogen.Header, ctx *jsonpointer.ResolveCtx) (_ map[string]*openapi.Header, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *parser) parseHeader(name string, header *ogen.Header, ctx *jsonpointer.ResolveCtx) (_ *openapi.Header, rerr error) {
	_ = "STUB: not implemented"
	return nil, nil
}
