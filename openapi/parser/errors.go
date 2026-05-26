package parser

import (
	"github.com/ogen-go/ogen/jsonpointer"
	"github.com/ogen-go/ogen/location"
)

// LocationError is a wrapper for an error that has a location.
type LocationError = location.Error

func (p *parser) file(ctx *jsonpointer.ResolveCtx) location.File {
	_ = "STUB: not implemented"
	return *new(location.File)
}

func (p *parser) wrapRef(file location.File, l location.Locator, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) wrapField(field string, file location.File, l location.Locator, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) wrapLocation(file location.File, l location.Locator, err error) error {
	_ = "STUB: not implemented"
	return nil

	// Do not wrap error if it is nil or is already a LocationError.
}
