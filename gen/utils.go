package gen

import (
	"go.uber.org/zap"

	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/jsonschema"
	"github.com/ogen-go/ogen/location"
)

func unreachable(v any) string { _ = "STUB: not implemented"; return "" }

func isBinary(s *jsonschema.Schema) bool { _ = "STUB: not implemented"; return false }

func isStream(s *jsonschema.Schema) bool {
	_ = "STUB: not implemented"
	// https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md#considerations-for-file-uploads
	//
	// The Spec says:
	//
	//	Content transferred in binary (octet-stream) MAY omit schema.
	return false
}

// Allow format to be empty, stream body often defined as just string.

// TODO(tdakkota): check ContentEncoding field

// isMultipartFile tries to map field to multipart file.
//
// Returns nil type if field is not a file parameter.
func isMultipartFile(ctx *genctx, t *ir.Type, p *jsonschema.Property) (*ir.Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func statusText(code int) string { _ = "STUB: not implemented"; return "" }

type position interface {
	Position() (location.Position, bool)
	File() location.File
}

func zapPosition(l position) zap.Field { _ = "STUB: not implemented"; return *new(zap.Field) }
