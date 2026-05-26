package jsonschema

import (
	"context"
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

// ExternalResolver resolves external links.
type ExternalResolver interface {
	Get(ctx context.Context, loc string) ([]byte, error)
}

var _ ExternalResolver = NoExternal{}

// NoExternal is ExternalResolver that always returns error.
type NoExternal struct{}

// Get implements ExternalResolver.
func (n NoExternal) Get(context.Context, string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExternalOptions is external reference resolver options.
type ExternalOptions struct {
	// HTTPClient sets http client to use. Defaults to http.DefaultClient.
	HTTPClient *http.Client

	// ReadFile sets function for reading files from fs. Defaults to os.ReadFile.
	ReadFile func(p string) ([]byte, error)
	// URLToFilePath sets function for converting url to file path. Defaults to urlpath.URLToFilePath.
	URLToFilePath func(u *url.URL) (string, error)

	// Logger sets logger to use. Defaults to zap.NewNop().
	Logger *zap.Logger
}

func (r *ExternalOptions) setDefaults() { _ = "STUB: not implemented"; return }

var _ ExternalResolver = externalResolver{}

type externalResolver struct {
	client        *http.Client
	readFile      func(p string) ([]byte, error)
	urlToFilePath func(u *url.URL) (string, error)
	logger        *zap.Logger
}

// NewExternalResolver creates new ExternalResolver.
//
// Currently only http(s) and file schemes are supported.
func NewExternalResolver(opts ExternalOptions) ExternalResolver {
	_ = "STUB: not implemented"
	return *new(ExternalResolver)
}

func (e externalResolver) httpGet(ctx context.Context, u *url.URL) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//#nosec G704

func (e externalResolver) Get(ctx context.Context, loc string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
