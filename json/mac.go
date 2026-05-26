package json

import (
	"net"

	"github.com/go-faster/jx"
)

// DecodeMAC decodes net.HardwareAddr.
func DecodeMAC(d *jx.Decoder) (net.HardwareAddr, error) {
	_ = "STUB: not implemented"
	return *new(net.HardwareAddr), nil
}

// EncodeMAC encodes net.HardwareAddr.
func EncodeMAC(e *jx.Encoder, v net.HardwareAddr) { _ = "STUB: not implemented"; return }
