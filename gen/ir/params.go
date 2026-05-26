package ir

import (
	"github.com/ogen-go/ogen/openapi"
)

func (op *Operation) PathParams() []*Parameter   { _ = "STUB: not implemented"; return nil }
func (op *Operation) QueryParams() []*Parameter  { _ = "STUB: not implemented"; return nil }
func (op *Operation) CookieParams() []*Parameter { _ = "STUB: not implemented"; return nil }
func (op *Operation) HeaderParams() []*Parameter { _ = "STUB: not implemented"; return nil }

func (op Operation) HasQueryParams() bool { _ = "STUB: not implemented"; return false }

func (op Operation) HasHeaderParams() bool { _ = "STUB: not implemented"; return false }

func (op Operation) HasCookieParams() bool { _ = "STUB: not implemented"; return false }

func (op Operation) PathParamsCount() (r int) { _ = "STUB: not implemented"; return 0 }

func (op Operation) PathParamIndex(name string) int { _ = "STUB: not implemented"; return 0 }

// Cut brackets '{', '}'.

func (op *Operation) getParams(locatedIn openapi.ParameterLocation) []*Parameter {
	_ = "STUB: not implemented"
	return nil
}
