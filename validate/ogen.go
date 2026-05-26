package validate

import (
	"sync"
)

// OgenValidator is a function that performs custom validation.
type OgenValidator func(value any, params any) error

// OgenValidatorRegistry holds custom validators that can be registered and used for validation.
type OgenValidatorRegistry struct {
	mu         sync.RWMutex
	validators map[string]OgenValidator
}

// NewOgenValidatorRegistry creates a new OgenValidatorRegistry.
func NewOgenValidatorRegistry() *OgenValidatorRegistry { _ = "STUB: not implemented"; return nil }

func (r *OgenValidatorRegistry) Register(name string, validator OgenValidator) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *OgenValidatorRegistry) Get(name string) (OgenValidator, bool) {
	_ = "STUB: not implemented"
	return *new(OgenValidator), false
}

// Validate validates a value using the specified validator and parameters.
func (r *OgenValidatorRegistry) Validate(validatorName string, value, params any) error {
	_ = "STUB: not implemented"
	return nil
}

// Wrap in ValidationError if it's not already one

// ValidationError represents a validation error from a custom validator.
type ValidationError struct {
	ValidatorName string
	Value         any
	Params        any
	Message       string
}

// Error implements the error interface.
func (e *ValidationError) Error() string { _ = "STUB: not implemented"; return "" }

// Default global registry
var defaultRegistry = NewOgenValidatorRegistry()

// RegisterValidator registers a validator in the default global registry.
func RegisterValidator(name string, validator OgenValidator) error {
	_ = "STUB: not implemented"
	return nil
}

// GetValidator returns a validator from the default global registry.
func GetValidator(name string) (OgenValidator, bool) {
	_ = "STUB: not implemented"
	return *new(OgenValidator), false
}

// Ogen validates using the default global registry.
func Ogen(name string, value, params any) error { _ = "STUB: not implemented"; return nil }
