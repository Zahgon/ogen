package json

import (
	"net/url"

	"github.com/go-faster/jx"
)

// DecodeURI decodes url.URL from json.
func DecodeURI(i *jx.Decoder) (v url.URL, err error) {
	_ = "STUB: not implemented"
	return *new(url.URL), nil
}

// EncodeURI encodes url.URL to json.
func EncodeURI(s *jx.Encoder, v url.URL) { _ = "STUB: not implemented"; return }
