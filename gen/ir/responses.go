package ir

type ResponseInfo struct {
	Type           *Type
	Encoding       Encoding
	ContentType    ContentType
	StatusCode     int
	NoContent      bool
	WithStatusCode bool
	WithHeaders    bool
	JSONStreaming  bool
	RawResponse    bool
	OpenTelemetry  bool
	Headers        map[string]*Parameter
}

func (r ResponseInfo) ContentTypeHeader() string { _ = "STUB: not implemented"; return "" }

var corsSimpleResponseHeaders = map[string]struct{}{
	"Cache-Control":    {},
	"Content-Language": {},
	"Content-Length":   {},
	"Content-Type":     {},
	"Expires":          {},
	"Last-Modified":    {},
	"Pragma":           {},
}

func (r ResponseInfo) ExposeHeadersHeader() string { _ = "STUB: not implemented"; return "" }

func sortResponseInfos(result []ResponseInfo) { _ = "STUB: not implemented"; return }

// Default responses has zero status code.

func (op *Operation) ListResponseTypes(otel bool) []ResponseInfo {
	_ = "STUB: not implemented"
	return nil
}
