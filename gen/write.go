package gen

import (
	"bytes"
	"regexp"
	"sync"
	"text/template"

	"github.com/ogen-go/ogen/gen/ir"
)

type TemplateConfig struct {
	Package           string
	Operations        []*ir.Operation
	DefaultOperations []*ir.Operation
	OperationGroups   []*ir.OperationGroup
	Webhooks          []*ir.Operation
	Types             map[string]*ir.Type
	Interfaces        map[string]*ir.Type
	Error             *ir.Response
	ErrorType         *ir.Type
	Servers           ir.Servers
	Securities        map[string]*ir.Security
	Router            Router
	WebhookRouter     WebhookRouter
	Imports           map[string]string

	PathsClientEnabled        bool
	PathsServerEnabled        bool
	WebhookClientEnabled      bool
	WebhookServerEnabled      bool
	OpenTelemetryEnabled      bool
	SecurityReentrantEnabled  bool
	RequestOptionsEnabled     bool
	RequestValidationEnabled  bool
	ResponseValidationEnabled bool
	EditorsEnabled            bool

	skipTestRegex *regexp.Regexp
}

// AnyClientEnabled returns true, if webhooks or paths client is enabled.
func (t TemplateConfig) AnyClientEnabled() bool { _ = "STUB: not implemented"; return false }

// AnyServerEnabled returns true, if webhooks or paths server is enabled.
func (t TemplateConfig) AnyServerEnabled() bool { _ = "STUB: not implemented"; return false }

// AnyInstrumentable returns true, if OpenTelemetry integration enabled and there is client/server to instrument.
func (t TemplateConfig) AnyInstrumentable() bool { _ = "STUB: not implemented"; return false }

// ErrorGoType returns Go type of error.
func (t TemplateConfig) ErrorGoType() string { _ = "STUB: not implemented"; return "" }

// SkipTest returns true, if test should be skipped.
func (t TemplateConfig) SkipTest(typ *ir.Type) bool { _ = "STUB: not implemented"; return false }

func (t TemplateConfig) collectStrings(cb func(typ *ir.Type) []string) []string {
	_ = "STUB: not implemented"
	return nil
}

// RegexStrings returns slice of all unique regex validators.
func (t TemplateConfig) RegexStrings() []string { _ = "STUB: not implemented"; return nil }

// RatStrings returns slice of all unique big.Rat (multipleOf validation).
func (t TemplateConfig) RatStrings() []string { _ = "STUB: not implemented"; return nil }

// `RatString` return a string with integer value if denominator is 1.
//
// That makes string representation of `big.Rat` shorter and simpler.
// Also, it is better for executable size.

// FileSystem represents a directory of generated package.
type FileSystem interface {
	WriteFile(baseName string, source []byte) error
}

type writer struct {
	fs FileSystem
	t  *template.Template
}

// generatorBufSize is 1 MB, it's enough for most mid-size specs.
const generatorBufSize = 1024 * 1024

var bufPool = sync.Pool{
	New: func() any {
		var b bytes.Buffer
		b.Grow(generatorBufSize)
		b.Reset()
		return &b
	},
}

func getBuffer() *bytes.Buffer { _ = "STUB: not implemented"; return nil }

func putBuffer(b *bytes.Buffer) { _ = "STUB: not implemented"; return }

// Generate executes template to file using config.
func (w *writer) Generate(templateName, fileName string, cfg TemplateConfig) (rerr error) {
	_ = "STUB: not implemented"
	return nil
}

// WriteSource writes generated definitions to fs.
func (g *Generator) WriteSource(fs FileSystem, pkgName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Historically we separate interfaces from other types.
// This is done for backward compatibility.

// Unused for now.

// Generate Equal() and Hash() methods for complex uniqueItems validation

// Generate validateUnique[TypeName]() functions for runtime validation

func (g *Generator) hasAnyType(check func(t *ir.Type) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *Generator) hasDefaultFields() bool { _ = "STUB: not implemented"; return false }

func (g *Generator) hasJSON() bool { _ = "STUB: not implemented"; return false }

func (g *Generator) hasValidators() bool { _ = "STUB: not implemented"; return false }

func (g *Generator) hasParams() bool { _ = "STUB: not implemented"; return false }

func (g *Generator) hasURIObjectParams() bool { _ = "STUB: not implemented"; return false }
