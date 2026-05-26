package gen

import (
	"net/url"
	"regexp"

	"github.com/go-faster/yaml"
	"go.uber.org/zap"

	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/jsonschema"
	"github.com/ogen-go/ogen/location"
	"github.com/ogen-go/ogen/openapi"
)

// Options is Generator options.
type Options struct {
	// Parser sets parser options.
	Parser ParseOptions `json:"parser" yaml:"parser"`

	// Generator sets generator options.
	Generator GenerateOptions `json:"generator" yaml:"generator"`

	// ExpandSpec is a path to expanded spec.
	ExpandSpec string `json:"expand" yaml:"expand"`

	// Logger to use.
	Logger *zap.Logger `json:"-" yaml:"-"`
}

// SetLocation sets File, RootURL and RemoteOptions using given path or URL
// and returns file data.
func (o *Options) SetLocation(p string, opts RemoteOptions) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *Options) setDefaults() { _ = "STUB: not implemented"; return }

// RemoteOptions is remote reference resolver options.
type RemoteOptions = jsonschema.ExternalOptions

// ParseOptions sets parsing options.
type ParseOptions struct {
	// InferSchemaType enables type inference for schemas. Schema parser will try to detect schema type
	// by its properties.
	InferSchemaType bool `json:"infer_types" yaml:"infer_types"`
	// AllowRemote enables remote references resolving.
	//
	// See https://github.com/ogen-go/ogen/issues/385.
	AllowRemote bool `json:"allow_remote" yaml:"allow_remote"`
	// RootURL is root URL for remote references resolving.
	RootURL *url.URL `json:"-" yaml:"-"`
	// Remote is remote reference resolver options.
	Remote RemoteOptions `json:"-" yaml:"-"`
	// SchemaDepthLimit is maximum depth of schema generation. Default is 1000.
	SchemaDepthLimit int `json:"depth_limit" yaml:"depth_limit"`
	// AuthenticationSchemes is the list of allowed HTTP Authorization schemes in a Security Scheme Object.
	// Default is the list defined in https://www.iana.org/assignments/http-authschemes/http-authschemes.xhtml.
	AuthenticationSchemes []string `json:"authentication_schemes" yaml:"authentication_schemes"`
	// AllowCrossTypeConstraints enables interpretation of cross-type schema constraints.
	// When true (default), constraints like pattern on numbers or maximum on strings
	// are interpreted and enforced via generated validation code.
	// Set to false for strict JSON Schema validation that rejects such constraints.
	// Default: true
	AllowCrossTypeConstraints *bool `json:"allow_cross_type_constraints,omitempty" yaml:"allow_cross_type_constraints,omitempty"`
	// DisallowDuplicateMethodPaths controls whether paths that normalize to the same
	// structure (e.g., /pets/{petId} and /pets/{id}) are allowed when they have
	// different HTTP methods.
	//
	// When false (default), paths with different parameter names but different HTTP methods
	// are allowed, and operations are disambiguated by path + params + method.
	//
	// When true, duplicate paths are always rejected per strict OpenAPI spec interpretation.
	DisallowDuplicateMethodPaths bool `json:"disallow_duplicate_method_paths" yaml:"disallow_duplicate_method_paths"`
	// File is the file that is being parsed.
	//
	// Used for error messages.
	File location.File `json:"-" yaml:"-"`
}

// SetLocation sets File, RootURL and RemoteOptions using given path or URL
// and returns file data.
func (o *ParseOptions) SetLocation(p string, opts RemoteOptions) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FIXME(tdakkota): pass context.

// Guard against reading local files in remote mode.

func (o *ParseOptions) setDefaults() { _ = "STUB: not implemented"; return }

// GenerateOptions sets generator options.
type GenerateOptions struct {
	// Features sets generator features.
	Features *FeatureOptions `json:"features" yaml:"features"`

	// Filters contains filters to skip operations.
	Filters Filters `json:"filters" yaml:"filters"`

	// IgnoreNotImplemented contains ErrNotImplemented messages to ignore.
	IgnoreNotImplemented []string `json:"ignore_not_implemented" yaml:"ignore_not_implemented"`
	// NotImplementedHook is hook for ErrNotImplemented errors.
	NotImplementedHook func(name string, err error) `json:"-" yaml:"-"`

	// ConvenientErrors control Convenient Errors feature.
	//
	// Default value is `auto` (0), NewError handler will be generated if possible.
	//
	// If value > 0 forces feature. An error will be returned if generator is unable to find common error pattern.
	//
	// If value < 0 disables feature entirely.
	ConvenientErrors ConvenientErrors `json:"convenient_errors" yaml:"convenient_errors"`
	// ContentTypeAliases contains content type aliases.
	ContentTypeAliases ContentTypeAliases `json:"content_type_aliases" yaml:"content_type_aliases"`
	// WildcardContentTypeDefault specifies the default encoding to use for wildcard
	// content types (*/* or application/*) when the schema is not binary.
	//
	// Common values: "application/json", "text/plain"
	//
	// If empty, wildcard content types are treated as unsupported and will cause
	// an error unless explicitly mapped via ContentTypeAliases.
	WildcardContentTypeDefault ir.Encoding `json:"wildcard_content_type_default" yaml:"wildcard_content_type_default"`
}

// ConvenientErrors is an option type to control `Convenient Errors` feature.
type ConvenientErrors int

// IsDisabled whether Convenient Errors is disabled.
func (c ConvenientErrors) IsDisabled() bool {
	_ = "STUB: not implemented"

	// IsForced whether Convenient Errors is forced.
	return false
}

func (c ConvenientErrors) IsForced() bool {
	_ = "STUB: not implemented"

	// String implements fmt.Stringer.
	return false
}

func (c ConvenientErrors) String() string { _ = "STUB: not implemented"; return "" }

// IsBoolFlag implements flag.boolFlag.
func (c *ConvenientErrors) IsBoolFlag() bool {
	_ = "STUB: not implemented"

	// UnmarshalYAML implements [yaml.Unmarshaler].
	return false
}

func (c *ConvenientErrors) UnmarshalYAML(n *yaml.Node) error { _ = "STUB: not implemented"; return nil }

// Set implements flag.Value.
func (c *ConvenientErrors) Set(s string) error { _ = "STUB: not implemented"; return nil }

// ContentTypeAliases maps content type to concrete ogen encoding.
type ContentTypeAliases map[string]ir.Encoding

// String implements fmt.Stringer.
func (m ContentTypeAliases) String() string { _ = "STUB: not implemented"; return "" }

// Set implements flag.Value.
func (m *ContentTypeAliases) Set(value string) error { _ = "STUB: not implemented"; return nil }

// Filters contains filters to skip operations.
type Filters struct {
	PathRegex *regexp.Regexp
	Methods   []string
}

// UnmarshalYAML implements [yaml.Unmarshaler].
func (f *Filters) UnmarshalYAML(n *yaml.Node) error { _ = "STUB: not implemented"; return nil }

func (f Filters) accept(op *openapi.Operation) bool { _ = "STUB: not implemented"; return false }
