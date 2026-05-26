package json

import (
	"github.com/go-faster/jx"
	"golang.org/x/exp/constraints"
)

func encodeStringUint[T constraints.Unsigned](e *jx.Encoder, v T) {
	_ = "STUB: not implemented"
	return
}

// Write first quote

// Write integer

// Write second quote

// EncodeStringUint encodes string uint to json.
func EncodeStringUint(e *jx.Encoder, v uint) { _ = "STUB: not implemented"; return }

// DecodeStringUint decodes string int from json.
func DecodeStringUint(d *jx.Decoder) (uint, error) { _ = "STUB: not implemented"; return 0, nil }

// EncodeStringUint8 encodes string uint8 to json.
func EncodeStringUint8(e *jx.Encoder, v uint8) { _ = "STUB: not implemented"; return }

// DecodeStringUint8 decodes string int8 from json.
func DecodeStringUint8(d *jx.Decoder) (uint8, error) { _ = "STUB: not implemented"; return 0, nil }

// EncodeStringUint16 encodes string uint16 to json.
func EncodeStringUint16(e *jx.Encoder, v uint16) { _ = "STUB: not implemented"; return }

// DecodeStringUint16 decodes string int16 from json.
func DecodeStringUint16(d *jx.Decoder) (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

// EncodeStringUint32 encodes string uint32 to json.
func EncodeStringUint32(e *jx.Encoder, v uint32) { _ = "STUB: not implemented"; return }

// DecodeStringUint32 decodes string int32 from json.
func DecodeStringUint32(d *jx.Decoder) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

// EncodeStringUint64 encodes string uint64 to json.
func EncodeStringUint64(e *jx.Encoder, v uint64) { _ = "STUB: not implemented"; return }

// DecodeStringUint64 decodes string int64 from json.
func DecodeStringUint64(d *jx.Decoder) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }
