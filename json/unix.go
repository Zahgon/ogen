package json

import (
	"time"

	"github.com/go-faster/jx"
)

// DecodeUnixSeconds decodes unix-seconds from json string.
func DecodeUnixSeconds(d *jx.Decoder) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// EncodeUnixSeconds encodes unix-seconds to json string.
func EncodeUnixSeconds(e *jx.Encoder, v time.Time) {
	_ = "STUB: not implemented"

	// DecodeUnixNano decodes unix-nano from json string.
	return
}

func DecodeUnixNano(d *jx.Decoder) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// EncodeUnixNano encodes unix-nano to json string.
func EncodeUnixNano(e *jx.Encoder, v time.Time) { _ = "STUB: not implemented"; return }

// DecodeUnixMicro decodes unix-micro from json string.
func DecodeUnixMicro(d *jx.Decoder) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// EncodeUnixMicro encodes unix-micro to json string.
func EncodeUnixMicro(e *jx.Encoder, v time.Time) { _ = "STUB: not implemented"; return }

// DecodeUnixMilli decodes unix-milli from json string.
func DecodeUnixMilli(d *jx.Decoder) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// EncodeUnixMilli encodes unix-milli to json string.
func EncodeUnixMilli(e *jx.Encoder, v time.Time) { _ = "STUB: not implemented"; return }
