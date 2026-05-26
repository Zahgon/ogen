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

func IntToString(v int) string     { _ = "STUB: not implemented"; return "" }
func Int8ToString(v int8) string   { _ = "STUB: not implemented"; return "" }
func Int16ToString(v int16) string { _ = "STUB: not implemented"; return "" }
func Int32ToString(v int32) string { _ = "STUB: not implemented"; return "" }
func Int64ToString(v int64) string { _ = "STUB: not implemented"; return "" }

func UintToString(v uint) string     { _ = "STUB: not implemented"; return "" }
func Uint8ToString(v uint8) string   { _ = "STUB: not implemented"; return "" }
func Uint16ToString(v uint16) string { _ = "STUB: not implemented"; return "" }
func Uint32ToString(v uint32) string { _ = "STUB: not implemented"; return "" }
func Uint64ToString(v uint64) string { _ = "STUB: not implemented"; return "" }

func Float32ToString(v float32) string { _ = "STUB: not implemented"; return "" }
func Float64ToString(v float64) string { _ = "STUB: not implemented"; return "" }

func DecimalToString(v decimal.Decimal) string { _ = "STUB: not implemented"; return "" }

func BoolToString(v bool) string { _ = "STUB: not implemented"; return "" }

func StringToString(v string) string { _ = "STUB: not implemented"; return "" }
func BytesToString(v []byte) string  { _ = "STUB: not implemented"; return "" }

//nolint:gosec // Unsafe conversion is intended for performance.

func TimeToString(v time.Time) string     { _ = "STUB: not implemented"; return "" }
func DateToString(v time.Time) string     { _ = "STUB: not implemented"; return "" }
func DateTimeToString(v time.Time) string { _ = "STUB: not implemented"; return "" }
func HTTPDateToString(v time.Time) string { _ = "STUB: not implemented"; return "" }

func UnixSecondsToString(v time.Time) string { _ = "STUB: not implemented"; return "" }
func UnixNanoToString(v time.Time) string    { _ = "STUB: not implemented"; return "" }
func UnixMicroToString(v time.Time) string   { _ = "STUB: not implemented"; return "" }
func UnixMilliToString(v time.Time) string   { _ = "STUB: not implemented"; return "" }

func DurationToString(v time.Duration) string { _ = "STUB: not implemented"; return "" }

func UUIDToString(v uuid.UUID) string { _ = "STUB: not implemented"; return "" }

func MACToString(v net.HardwareAddr) string { _ = "STUB: not implemented"; return "" }

func AddrToString(v netip.Addr) string { _ = "STUB: not implemented"; return "" }

func URLToString(v url.URL) string { _ = "STUB: not implemented"; return "" }

func StringIntToString(v int) string     { _ = "STUB: not implemented"; return "" }
func StringInt8ToString(v int8) string   { _ = "STUB: not implemented"; return "" }
func StringInt16ToString(v int16) string { _ = "STUB: not implemented"; return "" }
func StringInt32ToString(v int32) string { _ = "STUB: not implemented"; return "" }
func StringInt64ToString(v int64) string { _ = "STUB: not implemented"; return "" }

func StringUintToString(v uint) string     { _ = "STUB: not implemented"; return "" }
func StringUint8ToString(v uint8) string   { _ = "STUB: not implemented"; return "" }
func StringUint16ToString(v uint16) string { _ = "STUB: not implemented"; return "" }
func StringUint32ToString(v uint32) string { _ = "STUB: not implemented"; return "" }
func StringUint64ToString(v uint64) string { _ = "STUB: not implemented"; return "" }

func StringFloat32ToString(v float32) string { _ = "STUB: not implemented"; return "" }
func StringFloat64ToString(v float64) string { _ = "STUB: not implemented"; return "" }

func StringDecimalToString(v decimal.Decimal) string { _ = "STUB: not implemented"; return "" }

type (
	marshaler[T any] interface {
		json.Marshaler
		*T
	}
	textMarshaler[T any] interface {
		encoding.TextMarshaler
		*T
	}
	binaryMarshaler[T any] interface {
		encoding.BinaryMarshaler
		*T
	}
	jsonMarshaler[T any] interface {
		stdjson.Marshaler
		*T
	}
)

func NativeToString[T any, P marshaler[T]](v T) string { _ = "STUB: not implemented"; return "" }

func StringNativeToString[T any, P marshaler[T]](v T) string { _ = "STUB: not implemented"; return "" }

func TextToString[T any, P textMarshaler[T]](v T) string { _ = "STUB: not implemented"; return "" }

func BinaryToString[T any, P binaryMarshaler[T]](v T) string { _ = "STUB: not implemented"; return "" }

func JSONToString[T any, P jsonMarshaler[T]](v T) string { _ = "STUB: not implemented"; return "" }

func StringJSONToString[T any, P jsonMarshaler[T]](v T) string {
	_ = "STUB: not implemented"
	return ""
}

func ExternalToString[T any](v T) string { _ = "STUB: not implemented"; return "" }

func StringExternalToString[T any](v T) string { _ = "STUB: not implemented"; return "" }

func encodeArray[T any](vs []T, encode func(T) string) []string {
	_ = "STUB: not implemented"
	return nil
}

func Int32ArrayToString(vs []int32) []string { _ = "STUB: not implemented"; return nil }

func Int64ArrayToString(vs []int64) []string { _ = "STUB: not implemented"; return nil }

func Float32ArrayToString(vs []float32) []string { _ = "STUB: not implemented"; return nil }

func Float64ArrayToString(vs []float64) []string { _ = "STUB: not implemented"; return nil }

func StringArrayToString(vs []string) []string { _ = "STUB: not implemented"; return nil }

func BytesArrayToString(vs [][]byte) []string { _ = "STUB: not implemented"; return nil }

func TimeArrayToString(vs []time.Time) []string { _ = "STUB: not implemented"; return nil }

func BoolArrayToString(vs []bool) []string { _ = "STUB: not implemented"; return nil }

func UUIDArrayToString(vs []uuid.UUID) []string { _ = "STUB: not implemented"; return nil }

func MACArrayToString(vs []net.HardwareAddr) []string { _ = "STUB: not implemented"; return nil }
