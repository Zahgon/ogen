package conv

import (
	"encoding"
	stdjson "encoding/json"
	"net"
	"net/netip"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/ogen-go/ogen/json"
)

func ToInt(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func ToInt8(s string) (int8, error) { _ = "STUB: not implemented"; return 0, nil }

func ToInt16(s string) (int16, error) { _ = "STUB: not implemented"; return 0, nil }

func ToInt32(s string) (int32, error) { _ = "STUB: not implemented"; return 0, nil }

func ToInt64(s string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func ToUint(s string) (uint, error) { _ = "STUB: not implemented"; return 0, nil }

func ToUint8(s string) (uint8, error) { _ = "STUB: not implemented"; return 0, nil }

func ToUint16(s string) (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

func ToUint32(s string) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func ToUint64(s string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func ToFloat32(s string) (float32, error) { _ = "STUB: not implemented"; return 0, nil }

func ToFloat64(s string) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func ToDecimal(s string) (decimal.Decimal, error) {
	_ = "STUB: not implemented"
	return *new(decimal.Decimal), nil
}

func ToString(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func ToBytes(s string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func ToTime(s string) (time.Time, error) { _ = "STUB: not implemented"; return *new(time.Time), nil }

func ToDate(s string) (time.Time, error) { _ = "STUB: not implemented"; return *new(time.Time), nil }

func ToDateTime(s string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func ToHTTPDate(s string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func ToUnixSeconds(s string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func ToUnixNano(s string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func ToUnixMicro(s string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func ToUnixMilli(s string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func ToBool(s string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func ToUUID(s string) (uuid.UUID, error) { _ = "STUB: not implemented"; return *new(uuid.UUID), nil }

func ToMAC(s string) (net.HardwareAddr, error) {
	_ = "STUB: not implemented"
	return *new(net.HardwareAddr), nil
}

func ToAddr(s string) (netip.Addr, error) { _ = "STUB: not implemented"; return *new(netip.Addr), nil }

func ToURL(s string) (url.URL, error) { _ = "STUB: not implemented"; return *new(url.URL), nil }

func ToDuration(s string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func ToStringInt(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func ToStringInt8(s string) (int8, error) { _ = "STUB: not implemented"; return 0, nil }

func ToStringInt16(s string) (int16, error) { _ = "STUB: not implemented"; return 0, nil }

func ToStringInt32(s string) (int32, error) { _ = "STUB: not implemented"; return 0, nil }

func ToStringInt64(s string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func ToStringUint(s string) (uint, error) { _ = "STUB: not implemented"; return 0, nil }

func ToStringUint8(s string) (uint8, error) { _ = "STUB: not implemented"; return 0, nil }

func ToStringUint16(s string) (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

func ToStringUint32(s string) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func ToStringUint64(s string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func ToStringFloat32(s string) (float32, error) { _ = "STUB: not implemented"; return 0, nil }

func ToStringFloat64(s string) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func ToStringDecimal(s string) (decimal.Decimal, error) {
	_ = "STUB: not implemented"
	return *new(decimal.Decimal), nil
}

type (
	ogenUnmarshaler[T any] interface {
		json.Unmarshaler
		*T
	}
	textUnmarshaler[T any] interface {
		encoding.TextUnmarshaler
		*T
	}
	binaryUnmarshaler[T any] interface {
		encoding.BinaryUnmarshaler
		*T
	}
	jsonUnmarshaler[T any] interface {
		stdjson.Unmarshaler
		*T
	}
)

// toBytes converts a string to a byte slice with zero allocation.
func toBytes(s string) []byte { _ = "STUB: not implemented"; return nil }

//nolint:gosec // Unsafe conversion is intended for performance.

func ToNative[T any, P ogenUnmarshaler[T]](s string) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func ToStringNative[T any, P ogenUnmarshaler[T]](s string) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func ToText[T any, P textUnmarshaler[T]](s string) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func ToBinary[T any, P binaryUnmarshaler[T]](s string) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func ToJSON[T any, P jsonUnmarshaler[T]](s string) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func ToStringJSON[T any, P jsonUnmarshaler[T]](s string) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func ToExternal[T any](s string) (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

func ToStringExternal[T any](s string) (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

func decodeArray[T any](a []string, decode func(string) (T, error)) ([]T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToInt32Array(a []string) ([]int32, error) { _ = "STUB: not implemented"; return nil, nil }

func ToInt64Array(a []string) ([]int64, error) { _ = "STUB: not implemented"; return nil, nil }

func ToFloat32Array(a []string) ([]float32, error) { _ = "STUB: not implemented"; return nil, nil }

func ToFloat64Array(a []string) ([]float64, error) { _ = "STUB: not implemented"; return nil, nil }

func ToStringArray(a []string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ToBytesArray(a []string) ([][]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func ToTimeArray(a []string) ([]time.Time, error) { _ = "STUB: not implemented"; return nil, nil }

func ToBoolArray(a []string) ([]bool, error) { _ = "STUB: not implemented"; return nil, nil }

func ToUUIDArray(a []string) ([]uuid.UUID, error) { _ = "STUB: not implemented"; return nil, nil }

func ToMACArray(a []string) ([]net.HardwareAddr, error) { _ = "STUB: not implemented"; return nil, nil }
