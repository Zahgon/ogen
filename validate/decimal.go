package validate

import (
	"github.com/shopspring/decimal"
)

// Decimal validates decimal numbers.
type Decimal struct {
	MultipleOf    decimal.Decimal
	MultipleOfSet bool

	Min          decimal.Decimal
	MinSet       bool
	MinExclusive bool

	Max          decimal.Decimal
	MaxSet       bool
	MaxExclusive bool
}

// SetMultipleOf sets multipleOf validator.
func (t *Decimal) SetMultipleOf(d decimal.Decimal) { _ = "STUB: not implemented"; return }

// SetExclusiveMinimum sets exclusive minimum value.
func (t *Decimal) SetExclusiveMinimum(v decimal.Decimal) { _ = "STUB: not implemented"; return }

// SetExclusiveMaximum sets exclusive maximum value.
func (t *Decimal) SetExclusiveMaximum(v decimal.Decimal) { _ = "STUB: not implemented"; return }

// SetMinimum sets minimum value.
func (t *Decimal) SetMinimum(v decimal.Decimal) { _ = "STUB: not implemented"; return }

// SetMaximum sets maximum value.
func (t *Decimal) SetMaximum(v decimal.Decimal) { _ = "STUB: not implemented"; return }

// Set reports whether any validations are set.
func (t Decimal) Set() bool { _ = "STUB: not implemented"; return false }

// Validate returns error if v does not match validation rules.
func (t Decimal) Validate(v decimal.Decimal) error { _ = "STUB: not implemented"; return nil }

func (t Decimal) validate(v decimal.Decimal) error { _ = "STUB: not implemented"; return nil }
