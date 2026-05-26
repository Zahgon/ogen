package ogen

import (
	"github.com/ogen-go/ogen/jsonschema"
)

// ToJSONSchema converts Schema to jsonschema.Schema.
func (s *Schema) ToJSONSchema() *jsonschema.RawSchema { _ = "STUB: not implemented"; return nil }

// ToJSONSchema converts Properties to jsonschema.RawProperties.
func (p Properties) ToJSONSchema() jsonschema.RawProperties {
	_ = "STUB: not implemented"
	return *new(jsonschema.RawProperties)
}

// ToJSONSchema converts Property to jsonschema.Property.
func (p Property) ToJSONSchema() jsonschema.RawProperty {
	_ = "STUB: not implemented"
	return *new(jsonschema.RawProperty)
}

// ToJSONSchema converts AdditionalProperties to jsonschema.AdditionalProperties.
func (p *AdditionalProperties) ToJSONSchema() *jsonschema.AdditionalProperties {
	_ = "STUB: not implemented"
	return nil
}

// ToJSONSchema converts PatternProperties to jsonschema.RawPatternProperties.
func (p PatternProperties) ToJSONSchema() (result jsonschema.RawPatternProperties) {
	_ = "STUB: not implemented"
	return *new(jsonschema.RawPatternProperties)
}

// ToJSONSchema converts Items to jsonschema.RawItems.
func (p *Items) ToJSONSchema() *jsonschema.RawItems { _ = "STUB: not implemented"; return nil }

// ToJSONSchema converts Discriminator to jsonschema.RawDiscriminator.
func (d *Discriminator) ToJSONSchema() *jsonschema.RawDiscriminator {
	_ = "STUB: not implemented"
	return nil
}

// ToJSONSchema converts XML to jsonschema.XML.
func (d *XML) ToJSONSchema() *jsonschema.XML { _ = "STUB: not implemented"; return nil }
