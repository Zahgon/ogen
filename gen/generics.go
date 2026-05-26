package gen

import (
	"github.com/ogen-go/ogen/gen/ir"
)

func checkStructRecursions(s *ir.Type) error { _ = "STUB: not implemented"; return nil }

// Required.

func boxType(t *ir.Type, v ir.GenericVariant) (*ir.Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Do not wrap if
//  * type is Any
//  * type is Stream
//  * type is not nullable and not optional

// Do not wrap if type is Null primitive and generic is nullable only.

// Using special case for array nil value if possible.

func genericPostfix(t *ir.Type) (string, error) { _ = "STUB: not implemented"; return "", nil }
