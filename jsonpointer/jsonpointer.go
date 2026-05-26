// Package jsonpointer contains RFC 6901 JSON Pointer implementation.
package jsonpointer

import (
	"strings"

	"github.com/go-faster/yaml"
)

// Resolve takes given pointer and returns byte slice of requested value if any.
// If value not found, returns NotFoundError.
func Resolve(ptr string, node *yaml.Node) (*yaml.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note that length is bigger than 1.

// Fast-path to not parse URL.

func find(ptr string, node *yaml.Node) (*yaml.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cut first /.

func findIdx(n *yaml.Node, part string) (result *yaml.Node, ok bool, _ error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func findKey(n *yaml.Node, part string) (*yaml.Node, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

var unescapeReplacer = strings.NewReplacer(
	"~1", "/",
	"~0", "~",
)

func unescape(part string) string {
	_ = "STUB: not implemented"
	// Replacer always creates new string, check that unescape is really necessary.
	return ""
}
