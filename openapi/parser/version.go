package parser

import (
	"github.com/ogen-go/ogen/openapi"
)

func (p *parser) parseVersion() (rerr error) { _ = "STUB: not implemented"; return nil }

// FeatureVersionError is an error that is returned when a feature is used
// that requires a newer version of OpenAPI.
type FeatureVersionError struct {
	Feature string
	Minimum openapi.Version
	Actual  openapi.Version
}

// Error implements error.
func (f *FeatureVersionError) Error() string { _ = "STUB: not implemented"; return "" }

func (p *parser) requireMinorVersion(feature string, minor int) error {
	_ = "STUB: not implemented"
	return nil
}
