package gen

import (
	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/jsonschema"
)

// genctx is a generation context.
type genctx struct {
	global *tstorage // readonly
	local  *tstorage
}

func (g *genctx) saveType(t *ir.Type) error { _ = "STUB: not implemented"; return nil }

func (g *genctx) saveRef(ref jsonschema.Ref, e ir.Encoding, t *ir.Type) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *genctx) lookupRef(ref jsonschema.Ref, e ir.Encoding) (*ir.Type, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (g *genctx) saveResponse(ref jsonschema.Ref, r *ir.Response) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *genctx) saveWType(parent, ref jsonschema.Ref, t *ir.Type) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *genctx) saveParameter(ref jsonschema.Ref, r *ir.Parameter) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *genctx) lookupResponse(ref jsonschema.Ref) (*ir.Response, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (g *genctx) lookupWType(parent, ref jsonschema.Ref) (*ir.Type, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (g *genctx) lookupParameter(ref jsonschema.Ref) (*ir.Parameter, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
