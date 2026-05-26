package json

import (
	"encoding"
	"encoding/json"

	"github.com/go-faster/jx"
)

type (
	marshaler[T any] interface {
		Marshaler
		*T
	}
	unmarshaler[T any] interface {
		Unmarshaler
		*T
	}
	textMarshaler[T any] interface {
		encoding.TextMarshaler
		*T
	}
	textUnmarshaler[T any] interface {
		encoding.TextUnmarshaler
		*T
	}
	binaryMarshaler[T any] interface {
		encoding.BinaryMarshaler
		*T
	}
	binaryUnmarshaler[T any] interface {
		encoding.BinaryUnmarshaler
		*T
	}
	jsonMarshaler[T any] interface {
		json.Marshaler
		*T
	}
	jsonUnmarshaler[T any] interface {
		json.Unmarshaler
		*T
	}
)

// EncodeNative encodes a value using [Marshaler] interface.
func EncodeNative[T any, P marshaler[T]](e *jx.Encoder, v T) {
	_ = "STUB: not implemented"

	// DecodeNative decodes a value using [Unmarshaler] interface.
	return
}

func DecodeNative[T any, P unmarshaler[T]](d *jx.Decoder) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// EncodeText encodes a value using [encoding.TextMarshaler] interface.
func EncodeText[T any, P textMarshaler[T]](e *jx.Encoder, v T) { _ = "STUB: not implemented"; return }

// DecodeText decodes a value using [encoding.TextUnmarshaler] interface.
func DecodeText[T any, P textUnmarshaler[T]](d *jx.Decoder) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// EncodeStringText encodes a string value using [encoding.TextMarshaler] interface.
func EncodeStringText[T any, P textMarshaler[T]](e *jx.Encoder, v T) {
	_ = "STUB: not implemented"
	return
}

// DecodeStringText decodes a string value using [encoding.TextUnmarshaler] interface.
func DecodeStringText[T any, P textUnmarshaler[T]](d *jx.Decoder) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// EncodeBinary encodes a value using [encoding.BinaryMarshaler] interface.
func EncodeBinary[T any, P binaryMarshaler[T]](e *jx.Encoder, v T) {
	_ = "STUB: not implemented"
	return
}

// DecodeBinary decodes a value using [encoding.BinaryUnmarshaler] interface.
func DecodeBinary[T any, P binaryUnmarshaler[T]](d *jx.Decoder) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// EncodeJSON encodes a value using [json.Marshaler] interface.
func EncodeJSON[T any, P jsonMarshaler[T]](e *jx.Encoder, v T) { _ = "STUB: not implemented"; return }

// DecodeJSON decodes a value using [json.Marshaler] interface.
func DecodeJSON[T any, P jsonUnmarshaler[T]](d *jx.Decoder) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// EncodeExternal encodes a value using [json.Marshal].
func EncodeExternal[T any](e *jx.Encoder, v T) { _ = "STUB: not implemented"; return }

// DecodeExternal decodes a value using [json.Unmarshal].
func DecodeExternal[T any](d *jx.Decoder) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}
