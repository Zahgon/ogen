package testtypes

import (
	"github.com/go-faster/jx"
)

type StringOgen struct{ Value string }

func (o *StringOgen) Encode(e *jx.Encoder) { _ = "STUB: not implemented"; return }

func (o *StringOgen) Decode(d *jx.Decoder) error { _ = "STUB: not implemented"; return nil }

type NumberOgen struct{ Value float64 }

func (o *NumberOgen) Encode(e *jx.Encoder) { _ = "STUB: not implemented"; return }

func (o *NumberOgen) Decode(d *jx.Decoder) error { _ = "STUB: not implemented"; return nil }

type StringJSON struct{ Value string }

func (j *StringJSON) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (j *StringJSON) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

type NumberJSON struct{ Value float64 }

func (j *NumberJSON) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (j *NumberJSON) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

type Text struct{ Value string }

func (t *Text) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Text) UnmarshalText(data []byte) error { _ = "STUB: not implemented"; return nil }

type Binary struct{ Value string }

func (b *Binary) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (b *Binary) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

type String string

type Number float64
