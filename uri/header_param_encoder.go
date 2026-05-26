package uri

import (
	"net/http"
)

type headerParamEncoder struct {
	*receiver
	paramName string
	explode   bool
	header    http.Header
}

func (e *headerParamEncoder) serialize() error { _ = "STUB: not implemented"; return nil }

// As per RFC6265:
//
// Origin servers SHOULD NOT fold multiple Set-Cookie header fields into
// a single header field. The usual mechanism for folding HTTP headers
// fields (i.e., as defined in RFC2616) might change the semantics of
// the Set-Cookie header field because the %x2C (",") character is used
// by Set-Cookie in a way that conflicts with such folding.
