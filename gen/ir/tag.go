package ir

import (
	"github.com/ogen-go/ogen/openapi"
)

// Tag of Field or Parameter.
type Tag struct {
	JSON      string             // json tag, empty for none
	Form      *openapi.Parameter // query form parameter
	ExtraTags map[string]string  // a map of extra struct field tags
}

// EscapedJSON returns quoted and escaped JSON tag.
func (t Tag) EscapedJSON() string { _ = "STUB: not implemented"; return "" }

// GetTags returns a formatted list of struct tags, which must be quoted by '`'
func (t Tag) GetTags() string { _ = "STUB: not implemented"; return "" }
