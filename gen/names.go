package gen

import (
	"iter"

	"github.com/ogen-go/ogen/jsonschema"
)

func cleanRef(r jsonschema.Ref) string { _ = "STUB: not implemented"; return "" }

// Cuts file name.
//
// https://example.com/foo/bar.json -> bar
// foo/bar.json -> bar

type nameGen struct {
	parts []string
	src   []rune
	pos   int

	allowSpecial bool // special characters like +, -, /
}

func (g *nameGen) next() (rune, bool) { _ = "STUB: not implemented"; return 0, false }

var namedChar = map[rune][]rune{
	'+': []rune("Plus"),
	'-': []rune("Minus"),
	'/': []rune("Slash"),
	'<': []rune("Less"),
	'>': []rune("Greater"),
	'=': []rune("Eq"),
	'.': []rune("Dot"),
}

func (g *nameGen) generate() (string, error) { _ = "STUB: not implemented"; return "", nil }

// FIXME(tdakkota): choose prefix according to context

func (g *nameGen) clean() string { _ = "STUB: not implemented"; return "" }

func (g *nameGen) isAllowed(r rune) bool { _ = "STUB: not implemented"; return false }

func (g *nameGen) checkPart(part string) string { _ = "STUB: not implemented"; return "" }

func cleanSpecial(strs ...string) string { _ = "STUB: not implemented"; return "" }

func pascal(strs ...string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func pascalSpecial(strs ...string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func pascalNonEmpty(strs ...string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func camel(s ...string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func camelSpecial(s ...string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// firstLower returns s with first rune mapped to lower case.
func firstLower(s string) string { _ = "STUB: not implemented"; return "" }

// valueMappingNameGen creates a name generator for either an enum or discriminator mapping
func valueMappingNameGen(
	mapType, name string,
	values iter.Seq[any],
	valuesLen int,
	allowSpecial bool,
) (func(v any, idx int) (string, error), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This code is called when vstrCache is fully populated, so it's ok.

// Do not use pascal strategy for enum values starting with special characters.
//
// This rule is created to be able to distinguish
// between negative and positive numbers in this case:
//
// enum:
//   - '1'
//   - '-2'
//   - '3'
//   - '-4'

// Treat enum type name as duplicate to prevent collisions.

// enumVariantNameGen creates a name generator for enum values.
func enumVariantNameGen(name string, values []any) (func(v any, idx int) (string, error), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// discriminatorMappingNameGen creates a name generator for discriminator mapping keys.
func discriminatorMappingNameGen(name string, keys []string) (func(v any, idx int) (string, error), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create sequence that yields the mapping keys

// ensure user_id => UserId and not UserID
