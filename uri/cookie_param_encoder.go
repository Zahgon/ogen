package uri

import (
	"net/http"
)

type cookieParamEncoder struct {
	*receiver
	paramName string
	explode   bool
	req       *http.Request
}

func (e *cookieParamEncoder) setCookie(val string) {
	_ = "STUB: not implemented"
	// #nosec G124
	return
}

func (e *cookieParamEncoder) serialize() error { _ = "STUB: not implemented"; return nil }
