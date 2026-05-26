package gen

import (
	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/jsonschema"
)

type schemaKey struct {
	jsonschema.Ref
	ir.Encoding
}

// tstorage is a type storage.
type tstorage struct {
	refs map[schemaKey]*ir.Type // Key: ref

	// types map contains public types.
	// Public type is any type that has a name:
	//  * Struct
	//  * Alias
	//  * Generic
	//  * Interface
	//  * etc
	//
	// Example:
	// ...
	// requestBody:
	//   content:
	//     application/json:
	//       schema:
	//         type: string <- this type will not present
	//                         in the map because
	//                         the type is anonymous.
	//
	types      map[string]*ir.Type              // Key: type name
	responses  map[jsonschema.Ref]*ir.Response  // Key: ref
	parameters map[jsonschema.Ref]*ir.Parameter // Key: ref

	// wtypes stores references to wrapped types:
	//  * [T]StatusCode
	//  * [T]Headers
	//  * [T]StatusCodeWithHeaders
	wtypes map[[2]jsonschema.Ref]*ir.Type // Key: parent ref + ref
}

func newTStorage() *tstorage { _ = "STUB: not implemented"; return nil }

func (s *tstorage) saveType(t *ir.Type) error { _ = "STUB: not implemented"; return nil }

// HACK:
// Currently generator can overwrite same generic type
// multiple times during IR generation.
//
// We need to keep the set of features and methods consistent
// during this overwrites...
//
// Maybe we should instantiate generic types only once when needed
// and reuse them?

func (s *tstorage) saveRef(ref jsonschema.Ref, e ir.Encoding, t *ir.Type) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *tstorage) saveResponse(ref jsonschema.Ref, r *ir.Response) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *tstorage) saveWType(parent, ref jsonschema.Ref, t *ir.Type) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *tstorage) saveParameter(ref jsonschema.Ref, p *ir.Parameter) error {
	_ = "STUB: not implemented"
	return nil
}

func sameBase(t, tt *ir.Type) bool { _ = "STUB: not implemented"; return false }

func (s *tstorage) merge(other *tstorage) error {
	_ = "STUB: not implemented"
	// Check for merge conflicts.
	return nil
}

// Merge types.
