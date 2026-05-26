package validate

import (
	"math/big"

	"github.com/ogen-go/ogen/ogenregex"
)

// Float validates float numbers.
type Float struct {
	MultipleOf    *big.Rat
	MultipleOfSet bool

	Min          float64
	MinSet       bool
	MinExclusive bool

	Max          float64
	MaxSet       bool
	MaxExclusive bool

	// Pattern constraint for validating string representation
	Pattern ogenregex.Regexp
}

// SetMultipleOf sets multipleOf validator.
func (t *Float) SetMultipleOf(rat *big.Rat) { _ = "STUB: not implemented"; return }

// SetExclusiveMinimum sets exclusive minimum value.
func (t *Float) SetExclusiveMinimum(v float64) { _ = "STUB: not implemented"; return }

// SetExclusiveMaximum sets exclusive maximum value.
func (t *Float) SetExclusiveMaximum(v float64) { _ = "STUB: not implemented"; return }

// SetMinimum sets minimum value.
func (t *Float) SetMinimum(v float64) { _ = "STUB: not implemented"; return }

// SetMaximum sets maximum value.
func (t *Float) SetMaximum(v float64) { _ = "STUB: not implemented"; return }

// SetPattern sets pattern constraint for validating string representation.
func (t *Float) SetPattern(v ogenregex.Regexp) {
	_ = "STUB: not implemented"

	// Set reports whether any validations are set.
	return
}

func (t Float) Set() bool { _ = "STUB: not implemented"; return false }

// Validate returns error if v does not match validation rules.
func (t Float) Validate(v float64) error { _ = "STUB: not implemented"; return nil }

// ValidateStringified returns error if v does not match validation rules.
func (t Float) ValidateStringified(v float64) error { _ = "STUB: not implemented"; return nil }

func (t Float) validate(v float64) error { _ = "STUB: not implemented"; return nil }

// Validate pattern on string representation
