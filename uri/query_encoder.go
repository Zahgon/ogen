package uri

import (
	"mime/multipart"
	"net/url"
	"strings"
)

type QueryEncoder struct {
	values url.Values

	ct map[string]string
}

func NewQueryEncoder() *QueryEncoder { _ = "STUB: not implemented"; return nil }

func NewFormEncoder(ct map[string]string) *QueryEncoder { _ = "STUB: not implemented"; return nil }

type QueryParameterEncodingConfig struct {
	Name    string
	Style   QueryStyle
	Explode bool
}

func (e *QueryEncoder) EncodeParam(cfg QueryParameterEncodingConfig, f func(Encoder) error) error {
	_ = "STUB: not implemented"
	return nil
}

// FIXME(tdakkota): probable we should return the error during encoding

func (e *QueryEncoder) Values() url.Values { _ = "STUB: not implemented"; return *new(url.Values) }

func (e *QueryEncoder) WriteMultipart(w *multipart.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *QueryEncoder) writeMultipartField(w *multipart.Writer, key, value string) error {
	_ = "STUB: not implemented"
	return nil
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string { _ = "STUB: not implemented"; return "" }
