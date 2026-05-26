package gen

import (
	"go.uber.org/zap"

	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/jsonschema"
	"github.com/ogen-go/ogen/openapi"
)

// normalizeContentEncoding parses the media type, applies aliases and
// normalizes +json suffix media types as JSON per RFC 6838.
// This allows types like application/json-patch+json,
// application/vnd.api+json, application/merge-patch+json, etc.
// to work without explicit ContentTypeAliases.
// Note: application/problem+json has special handling via EncodingProblemJSON.
func normalizeContentEncoding(contentType string, aliases ContentTypeAliases,
) (parsedContentType string, encoding ir.Encoding, err error) {
	_ = "STUB: not implemented"
	return "", *new(ir.Encoding), nil
}

func isJSONLikeEncoding(encoding ir.Encoding) bool { _ = "STUB: not implemented"; return false }

func filterMostSpecific(contents map[string]*openapi.MediaType, log *zap.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// Special case for "*", "**", etc.

// There is at least one another media type, so delete "*".

// There is no other media type, so keep "*".

// Do not try to match mask against itself.

// Found more specific media type that matches the mask, so delete the mask.

// Found no more specific media type, so keep the mask.

func (g *Generator) wrapContent(ctx *genctx, name string, t *ir.Type) (ret *ir.Type, rerr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) generateFormContent(
	ctx *genctx,
	typeName string,
	media *openapi.MediaType,
	optional bool,
	encoding ir.Encoding,
) (*ir.Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A funny moment when you have a spec that shares schema between multipart form and JSON request and
// at some point you made ingenious decision to keep all types in one package at the same time.

func isComplexMultipartType(s *jsonschema.Schema) bool { _ = "STUB: not implemented"; return false }

func (g *Generator) generateContents(
	ctx *genctx,
	name string,
	optional,
	request bool,
	contents map[string]*openapi.MediaType,
) (_ map[ir.ContentType]ir.Media, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Handle wildcard content types using configured default
// unless the schema is binary (which is already handled by isStream check below)

// Check if this is a binary stream - if so, keep default behavior

// Use the mapped content type for the result key to avoid wrapping

// In rfc9457, the only MUST defined for generators is to keep the status field
// synced with the HTTP status code.

// FIXME(tdakkota): box if optional is true?
