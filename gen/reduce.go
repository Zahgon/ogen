package gen

import (
	"github.com/ogen-go/ogen/jsonschema"
	"github.com/ogen-go/ogen/openapi"
)

// reduceDefault implements convenient errors, representing common default
// response as error instead of variant of each response.
func (g *Generator) reduceDefault(ops []*openapi.Operation) error {
	_ = "STUB: not implemented"
	return nil
}

// Compare first default response to others.
//
// TODO(tdakkota): reduce by 4XX/5XX?

// TODO(tdakkota): point to "content", not to the entire response

// TODO(tdakkota): point to "content", not to the entire response

type responseComparator struct{}

func (c responseComparator) compare(a, b *openapi.Response) bool {
	_ = "STUB: not implemented"
	// Compile time check to not forget to update compareResponses.
	return false
}

func (c responseComparator) compareHeader(a, b *openapi.Header) bool {
	_ = "STUB: not implemented"
	return false
}

func (c responseComparator) compareParameterContent(a, b *openapi.ParameterContent) bool {
	_ = "STUB: not implemented"
	return false
}

func (c responseComparator) compareMediaType(a, b *openapi.MediaType) bool {
	_ = "STUB: not implemented"
	return false
}

func (c responseComparator) compareEncoding(a, b *openapi.Encoding) bool {
	_ = "STUB: not implemented"
	return false
}

func (c responseComparator) compareSchema(a, b *jsonschema.Schema) bool {
	_ = "STUB: not implemented"
	return false
}

func (c responseComparator) comparePatternProperty(a, b jsonschema.PatternProperty) bool {
	_ = "STUB: not implemented"
	return false
}

func (c responseComparator) compareProperty(a, b jsonschema.Property) bool {
	_ = "STUB: not implemented"
	return false
}

func (c responseComparator) compareDiscriminator(a, b *jsonschema.Discriminator) bool {
	_ = "STUB: not implemented"
	return false
}

func (c responseComparator) compareXML(a, b *jsonschema.XML) bool {
	_ = "STUB: not implemented"
	return false
}

func (c responseComparator) compareNum(a, b jsonschema.Num) bool {
	_ = "STUB: not implemented"
	return false
}
