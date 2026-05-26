package parser

import (
	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/openapi"
)

func fromOpenapiInfo(info openapi.Info) ogen.Info {
	_ = "STUB: not implemented"
	return *new(ogen.Info)
}

func fromOgenInfo(info ogen.Info) openapi.Info {
	_ = "STUB: not implemented"
	return *new(openapi.Info)
}
