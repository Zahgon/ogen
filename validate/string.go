package validate

import (
	"github.com/ogen-go/ogen/ogenregex"
)

// String validator.
type String struct {
	MinLength    int
	MinLengthSet bool
	MaxLength    int
	MaxLengthSet bool
	Email        bool
	Regex        ogenregex.Regexp
	Hostname     bool

	// Numeric constraints for strings representing numbers
	MinNumeric    float64
	MinNumericSet bool
	MaxNumeric    float64
	MaxNumericSet bool
}

// SetMaxLength sets maximum string length (in Unicode code points).
func (t *String) SetMaxLength(v int) { _ = "STUB: not implemented"; return }

// SetMinLength sets minimum string length (in Unicode code points).
func (t *String) SetMinLength(v int) { _ = "STUB: not implemented"; return }

// SetMaximumNumeric sets maximum numeric value for numeric strings.
func (t *String) SetMaximumNumeric(v float64) { _ = "STUB: not implemented"; return }

// SetMinimumNumeric sets minimum numeric value for numeric strings.
func (t *String) SetMinimumNumeric(v float64) { _ = "STUB: not implemented"; return }

// Set reports whether any validations are set.
func (t String) Set() bool { _ = "STUB: not implemented"; return false }

func (t String) checkHostname(v string) error { _ = "STUB: not implemented"; return nil }

func (t String) checkEmail(v string) error {
	_ = "STUB: not implemented"
	// Pretty basic validation, but should work for most cases and is not
	// too strict to break things.
	//
	// Still better than obscure regex or std `mail.ParseAddress`.
	return nil
}

// Validate returns error if v does not match validation rules.
func (t String) Validate(v string) error { _ = "STUB: not implemented"; return nil }

// Validate numeric constraints on string values

func (t String) validateNumeric(v string) error {
	_ = "STUB: not implemented"
	// Parse string as float64
	return nil
}
