// Package ogenzap contains ogen logging utilities.
package ogenzap

import (
	"flag"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// DefaultColorFlag returns default color flag value.
//
// See https://no-color.org.
func DefaultColorFlag() bool { _ = "STUB: not implemented"; return false }

// By default, NewDevelopmentConfig uses stderr.

// Options is options for Create.
type Options struct {
	Level     zapcore.Level
	Verbose   bool
	Color     bool
	FnOptions []zap.Option
}

// RegisterFlags registers fields of Options as flags.
func (o *Options) RegisterFlags(set *flag.FlagSet) { _ = "STUB: not implemented"; return }

// Create creates new logger for ogen.
func Create(opts Options) (*zap.Logger, error) { _ = "STUB: not implemented"; return nil, nil }

// Set to noop if logging is not verbose.

// Disable stacktrace and caller.
