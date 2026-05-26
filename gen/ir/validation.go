package ir

import (
	"github.com/ogen-go/ogen/jsonschema"
	"github.com/ogen-go/ogen/validate"
)

type Validators struct {
	String  validate.String
	Int     validate.Int
	Float   validate.Float
	Decimal validate.Decimal
	Array   validate.Array
	Object  validate.Object
	// Ogen contains parameters for custom validation.
	Ogen map[string]any
}

func (v *Validators) SetString(schema *jsonschema.Schema) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Interpret numeric constraints on string type.
// This handles specs that use maximum/minimum on strings representing numbers.

func (v *Validators) SetInt(schema *jsonschema.Schema) error { _ = "STUB: not implemented"; return nil }

// Interpret pattern constraint on integer type.
// This validates the string representation of the integer.

func (v *Validators) SetFloat(schema *jsonschema.Schema) error {
	_ = "STUB: not implemented"
	return nil
}

// Interpret pattern constraint on float type.
// This validates the string representation of the float.

func (v *Validators) SetDecimal(schema *jsonschema.Schema) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validators) SetArray(schema *jsonschema.Schema) { _ = "STUB: not implemented"; return }

func (v *Validators) SetObject(schema *jsonschema.Schema) { _ = "STUB: not implemented"; return }

func (t *Type) NeedValidation() bool { _ = "STUB: not implemented"; return false }

func (v *Validators) SetOgenValidate(schema *jsonschema.Schema) { _ = "STUB: not implemented"; return }

func (t *Type) needValidation(path *walkpath) (result bool) {
	_ = "STUB: not implemented"
	return false
}

// NaN, Inf, float validators.

// FIXME(tdakkota): try to validate Any.
