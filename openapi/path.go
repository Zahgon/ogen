package openapi

// Path is an operation path.
type Path []PathPart[*Parameter]

// ID returns path, but without parameter names.
//
// For example, if path is "/users/{id}", ID returns "/users/{}".
func (p Path) ID() (path string) { _ = "STUB: not implemented"; return "" }

// String implements fmt.Stringer.
func (p Path) String() (path string) { _ = "STUB: not implemented"; return "" }

// PathPart is a part of an OpenAPI Operation Path.
type PathPart[P any] struct {
	Raw   string
	Param P
}

// IsParam returns true if part is a parameter.
func (p PathPart[P]) IsParam() bool { _ = "STUB: not implemented"; return false }
