package json

import (
	"github.com/go-faster/jx"
	"github.com/google/uuid"
)

// DecodeUUID decodes UUID from json.
func DecodeUUID(i *jx.Decoder) (v uuid.UUID, err error) {
	_ = "STUB: not implemented"
	return *new(uuid.UUID), nil
}

// EncodeUUID encodes UUID to json.
func EncodeUUID(s *jx.Encoder, v uuid.UUID) {
	_ = "STUB: not implemented"

	// Hexed length (16 * 2) + 4 hyphens
	return
}

func hexEncode(dst *[36]byte, v uuid.UUID) { _ = "STUB: not implemented"; return }
