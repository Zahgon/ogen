package json

import (
	"time"

	"github.com/go-faster/jx"
)

// DecodeStringUnixSeconds decodes unix-seconds from json string.
func DecodeStringUnixSeconds(d *jx.Decoder) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// EncodeStringUnixSeconds encodes unix-seconds to json string.
func EncodeStringUnixSeconds(e *jx.Encoder, v time.Time) { _ = "STUB: not implemented"; return }

// DecodeStringUnixNano decodes unix-nano from json string.
func DecodeStringUnixNano(d *jx.Decoder) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// EncodeStringUnixNano encodes unix-nano to json string.
func EncodeStringUnixNano(e *jx.Encoder, v time.Time) { _ = "STUB: not implemented"; return }

// DecodeStringUnixMicro decodes unix-micro from json string.
func DecodeStringUnixMicro(d *jx.Decoder) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// EncodeStringUnixMicro encodes unix-micro to json string.
func EncodeStringUnixMicro(e *jx.Encoder, v time.Time) { _ = "STUB: not implemented"; return }

// DecodeStringUnixMilli decodes unix-milli from json string.
func DecodeStringUnixMilli(d *jx.Decoder) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// EncodeStringUnixMilli encodes unix-milli to json string.
func EncodeStringUnixMilli(e *jx.Encoder, v time.Time) { _ = "STUB: not implemented"; return }
