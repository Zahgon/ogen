package json

import (
	"github.com/go-faster/jx"
	"github.com/shopspring/decimal"
)

// EncodeDecimal encodes decimal.Decimal to json.
func EncodeDecimal(e *jx.Encoder, v decimal.Decimal) { _ = "STUB: not implemented"; return }

// DecodeDecimal decodes decimal.Decimal from json.
func DecodeDecimal(d *jx.Decoder) (decimal.Decimal, error) {
	_ = "STUB: not implemented"
	return *new(decimal.Decimal), nil
}

// EncodeStringDecimal encodes decimal.Decimal to json string.
func EncodeStringDecimal(e *jx.Encoder, v decimal.Decimal) {
	_ = "STUB: not implemented"

	// DecodeStringDecimal decodes decimal.Decimal from json string.
	return
}

func DecodeStringDecimal(d *jx.Decoder) (decimal.Decimal, error) {
	_ = "STUB: not implemented"
	return *new(decimal.Decimal), nil
}
