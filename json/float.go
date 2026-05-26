package json

import (
	"github.com/go-faster/jx"
	"golang.org/x/exp/constraints"
)

func encodeStringFloat[T constraints.Float](e *jx.Encoder, v T, bitSize int) {
	_ = "STUB: not implemented"
	return
}

// Write first quote

// Write float

// Write second quote

// EncodeStringFloat32 encodes string float32 to json.
func EncodeStringFloat32(e *jx.Encoder, v float32) { _ = "STUB: not implemented"; return }

// DecodeStringFloat32 decodes string float32 from json.
func DecodeStringFloat32(d *jx.Decoder) (float32, error) { _ = "STUB: not implemented"; return 0, nil }

// EncodeStringFloat64 encodes string float64 to json.
func EncodeStringFloat64(e *jx.Encoder, v float64) { _ = "STUB: not implemented"; return }

// DecodeStringFloat64 decodes string float64 from json.
func DecodeStringFloat64(d *jx.Decoder) (float64, error) { _ = "STUB: not implemented"; return 0, nil }
