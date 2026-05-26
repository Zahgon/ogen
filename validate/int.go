package validate

import (
	"github.com/ogen-go/ogen/ogenregex"
)

// Int validates integers.
type Int struct {
	MultipleOf    uint64
	MultipleOfSet bool

	Min          int64
	MinSet       bool
	MinExclusive bool

	Max          int64
	MaxSet       bool
	MaxExclusive bool

	// Pattern constraint for validating string representation
	Pattern ogenregex.Regexp
}

// SetMultipleOf sets multipleOf validator.
func (t *Int) SetMultipleOf(v uint64) { _ = "STUB: not implemented"; return }

// SetExclusiveMinimum sets exclusive minimum value.
func (t *Int) SetExclusiveMinimum(v int64) { _ = "STUB: not implemented"; return }

// SetExclusiveMaximum sets exclusive maximum value.
func (t *Int) SetExclusiveMaximum(v int64) { _ = "STUB: not implemented"; return }

// SetMinimum sets minimum value.
func (t *Int) SetMinimum(v int64) { _ = "STUB: not implemented"; return }

// SetMaximum sets maximum value.
func (t *Int) SetMaximum(v int64) { _ = "STUB: not implemented"; return }

// SetPattern sets pattern constraint for validating string representation.
func (t *Int) SetPattern(v ogenregex.Regexp) {
	_ = "STUB: not implemented"

	// Set reports whether any validations are set.
	return
}

func (t Int) Set() bool { _ = "STUB: not implemented"; return false }

// Validate returns error if v does not match validation rules.
func (t Int) Validate(v int64) error { _ = "STUB: not implemented"; return nil }

// We don't care about sign when checking value using multipleOf.

// Validate pattern on string representation
