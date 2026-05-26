package json

import (
	"time"

	"github.com/go-faster/jx"
)

const (
	dateLayout     = "2006-01-02"
	timeLayout     = "15:04:05"
	httpDateLayout = "Mon, 02 Jan 2006 15:04:05 GMT"
)

// DecodeTimeFormat decodes date, time & date-time from json using a custom layout.
func DecodeTimeFormat(d *jx.Decoder, layout string) (v time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// EncodeTimeFormat encodes date, time & date-time to json using a custom layout.
func EncodeTimeFormat(e *jx.Encoder, v time.Time, layout string) { _ = "STUB: not implemented"; return }

// Allocate buf on stack, if we can.

// NewTimeDecoder returns a new time decoder using a custom layout.
func NewTimeDecoder(layout string) func(i *jx.Decoder) (time.Time, error) {
	_ = "STUB: not implemented"
	return nil
}

// NewTimeEncoder returns a new time encoder using a custom layout.
func NewTimeEncoder(layout string) func(e *jx.Encoder, v time.Time) {
	_ = "STUB: not implemented"
	return nil
}

// DecodeDate decodes date from json.
func DecodeDate(d *jx.Decoder) (v time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// EncodeDate encodes date to json.
func EncodeDate(e *jx.Encoder, v time.Time) { _ = "STUB: not implemented"; return }

// DecodeTime decodes time from json.
func DecodeTime(d *jx.Decoder) (v time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// EncodeTime encodes time to json.
func EncodeTime(e *jx.Encoder, v time.Time) { _ = "STUB: not implemented"; return }

// DecodeDateTime decodes date-time from json.
func DecodeDateTime(d *jx.Decoder) (v time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// EncodeDateTime encodes date-time to json.
func EncodeDateTime(e *jx.Encoder, v time.Time) { _ = "STUB: not implemented"; return }

// DecodeHTTPDate decodes http-date from json.
func DecodeHTTPDate(d *jx.Decoder) (v time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// EncodeHTTPDate encodes http-date to json.
func EncodeHTTPDate(e *jx.Encoder, v time.Time) { _ = "STUB: not implemented"; return }

// DecodeDuration decodes duration from json.
func DecodeDuration(d *jx.Decoder) (v time.Duration, err error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// EncodeDuration encodes duration to json.
func EncodeDuration(e *jx.Encoder, v time.Duration) { _ = "STUB: not implemented"; return }
