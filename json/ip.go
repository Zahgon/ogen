package json

import (
	"net/netip"

	"github.com/go-faster/jx"
)

func decodeIP(d *jx.Decoder, checkVersion func(addr netip.Addr) bool) (v netip.Addr, err error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}

// DecodeIP decodes netip.Addr.
func DecodeIP(d *jx.Decoder) (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *

	// DecodeIPv4 decodes netip.Addr.
	new(netip.Addr), nil
}

func DecodeIPv4(d *jx.Decoder) (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}

// DecodeIPv6 decodes netip.Addr.
func DecodeIPv6(d *jx.Decoder) (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}

// EncodeIP encodes netip.Addr.
func EncodeIP(s *jx.Encoder, v netip.Addr) { _ = "STUB: not implemented"; return }

// EncodeIPv4 encodes netip.Addr.
func EncodeIPv4(s *jx.Encoder, v netip.Addr) {
	_ = "STUB: not implemented"

	// EncodeIPv6 encodes netip.Addr.
	return
}

func EncodeIPv6(s *jx.Encoder, v netip.Addr) { _ = "STUB: not implemented"; return }
