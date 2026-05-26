// Package json contains helper functions for encoding and decoding JSON.
package json

import (
	"github.com/go-faster/jx"
)

// Marshal value to json.
func Marshal(val any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Unmarshal value from json.
		nil
}

func Unmarshal(data []byte, val any) error { _ = "STUB: not implemented"; return nil }

// Unmarshaler implements json reading.
type Unmarshaler interface {
	Decode(d *jx.Decoder) error
}

// Marshaler implements json writing.
type Marshaler interface {
	Encode(e *jx.Encoder)
}

// Value represents a json value.
type Value interface {
	Marshaler
	Unmarshaler
}

// Settable value can be set (present) or unset
// (i.e. not provided or undefined).
type Settable interface {
	IsSet() bool
}

// Resettable value can be unset.
type Resettable interface {
	Reset()
}

// Nullable can be nil (but defined) or not.
type Nullable interface {
	IsNil() bool
}

// Encode Marshaler to byte slice.
func Encode(m Marshaler) []byte { _ = "STUB: not implemented"; return nil }
