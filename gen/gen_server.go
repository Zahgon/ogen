package gen

import (
	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/openapi"
)

func (g *Generator) generateServer(s openapi.Server) (ir.Server, error) {
	_ = "STUB: not implemented"
	return *new(ir.Server), nil
}

// The server name is passed using ogen extension, so it guaranteed to be
// valid Go identifier, but we need to make pascal case anyway.
