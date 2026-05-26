package validate

// Object validates map length.
type Object struct {
	MinProperties    int
	MinPropertiesSet bool
	MaxProperties    int
	MaxPropertiesSet bool
	// TODO: add validate to gen
	MinLength    int
	MinLengthSet bool
	MaxLength    int
	MaxLengthSet bool
}

// SetMinProperties sets MinProperties validation.
func (o *Object) SetMinProperties(v int) { _ = "STUB: not implemented"; return }

// SetMaxProperties sets MaxProperties validation.
func (o *Object) SetMaxProperties(v int) { _ = "STUB: not implemented"; return }

// SetMinLength sets MinLength validation.
func (o *Object) SetMinLength(v int) { _ = "STUB: not implemented"; return }

// SetMaxLength sets MaxLength validation.
func (o *Object) SetMaxLength(v int) { _ = "STUB: not implemented"; return }

// Set reports whether any validations are seo.
func (o Object) Set() bool { _ = "STUB: not implemented"; return false }

// ValidateProperties returns error if object length (properties number) v is invalid.
func (o Object) ValidateProperties(v int) error { _ = "STUB: not implemented"; return nil }
