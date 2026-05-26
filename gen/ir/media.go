package ir

// ContentType is a Content-Type header value.
type ContentType string

func (t ContentType) Mask() bool { _ = "STUB: not implemented"; return false }

func (t ContentType) String() string {
	_ = "STUB: not implemented"

	// Encoding of body.
	return ""
}

type Encoding string

const (
	// EncodingJSON is Encoding for json.
	EncodingJSON        Encoding = "application/json"
	EncodingProblemJSON Encoding = "application/problem+json"
	// EncodingFormURLEncoded is Encoding for URL-encoded form.
	EncodingFormURLEncoded Encoding = "application/x-www-form-urlencoded"
	// EncodingMultipart is Encoding for multipart form.
	EncodingMultipart Encoding = "multipart/form-data"
	// EncodingOctetStream is Encoding for binary.
	EncodingOctetStream Encoding = "application/octet-stream"
	// EncodingTextPlain is Encoding for text.
	EncodingTextPlain Encoding = "text/plain"
)

func (t Encoding) String() string { _ = "STUB: not implemented"; return "" }

func (t Encoding) JSON() bool { _ = "STUB: not implemented"; return false }

func (t Encoding) ProblemJSON() bool { _ = "STUB: not implemented"; return false }

func (t Encoding) FormURLEncoded() bool { _ = "STUB: not implemented"; return false }

func (t Encoding) MultipartForm() bool { _ = "STUB: not implemented"; return false }

func (t Encoding) OctetStream() bool { _ = "STUB: not implemented"; return false }

func (t Encoding) TextPlain() bool { _ = "STUB: not implemented"; return false }

type Media struct {
	// Encoding is the parsed content type used for encoding, but not for header value.
	Encoding Encoding
	// Type is response or request type.
	Type *Type

	// JSONStreaming indicates that the JSON media should be streamed.
	JSONStreaming bool
	// RawResponse indicates that the raw HTTP response should be returned.
	RawResponse bool
}
