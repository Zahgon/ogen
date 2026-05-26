package ir

import (
	"github.com/ogen-go/ogen/openapi"
)

// Servers is a list of servers.
type Servers []Server

func (s Servers) filter(cb func(Server) bool) (r []Server) { _ = "STUB: not implemented"; return nil }

// Templates returns a list of server URL templates.
func (s Servers) Templates() []Server { _ = "STUB: not implemented"; return nil }

// Const return a list of constant server URLs.
func (s Servers) Const() []Server { _ = "STUB: not implemented"; return nil }

// Server describes a OpenAPI server.
type Server struct {
	Name   string
	Params []ServerParam
	Spec   openapi.Server
}

// IsTemplate returns true if server URL has variables.
func (s Server) IsTemplate() bool { _ = "STUB: not implemented"; return false }

// ServerParam describes a server template parameter.
type ServerParam struct {
	// Name is a Go name of the parameter.
	Name string
	Spec openapi.ServerVariable
}

// FormatString returns a format string (fmt.Sprintf) for the server.
//
// If the server has no variables, returns plain string.
func (s Server) FormatString() string { _ = "STUB: not implemented"; return "" }

// GoDoc returns GoDoc comment for the server.
func (s Server) GoDoc() []string { _ = "STUB: not implemented"; return nil }
