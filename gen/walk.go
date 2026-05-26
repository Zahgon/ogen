package gen

import (
	"github.com/ogen-go/ogen/gen/ir"
)

func walkResponseTypes(r *ir.Responses, walkFn func(name string, t *ir.Type) (*ir.Type, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func walkOpTypes(ops []*ir.Operation, walk func(*ir.Type) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}
