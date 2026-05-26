package validate

// Array validates array length.
type Array struct {
	MinLength    int
	MinLengthSet bool
	MaxLength    int
	MaxLengthSet bool
	UniqueItems  bool
}

// SetMaxLength sets MaxLength validation.
func (t *Array) SetMaxLength(v int) { _ = "STUB: not implemented"; return }

// SetMinLength sets MinLength validation.
func (t *Array) SetMinLength(v int) { _ = "STUB: not implemented"; return }

// SetUniqueItems sets UniqueItems validation.
func (t *Array) SetUniqueItems(v bool) {
	_ = "STUB: not implemented"

	// Set reports whether any validations are set.
	return
}

func (t Array) Set() bool { _ = "STUB: not implemented"; return false }

// ValidateLength returns error if array length v is invalid.
func (t Array) ValidateLength(v int) error { _ = "STUB: not implemented"; return nil }

// UniqueItems ensures given array has no duplicates.
func UniqueItems[S ~[]T, T comparable](arr S) error { _ = "STUB: not implemented"; return nil }
