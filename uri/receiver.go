package uri

func checkNotContains(s, chars string) error { _ = "STUB: not implemented"; return nil }

type valueType string

func (v valueType) String() string { _ = "STUB: not implemented"; return "" }

const (
	typeNotSet valueType = "notSet"
	typeValue  valueType = "value"
	typeArray  valueType = "array"
	typeObject valueType = "object"
)

var _ Encoder = (*receiver)(nil)

// receiver is used to receive data from code generated types.
type receiver struct {
	typ    valueType
	val    string   // value type
	items  []string // array type
	fields []Field  // object type
}

func newReceiver() *receiver { _ = "STUB: not implemented"; return nil }

func (s *receiver) EncodeValue(v string) error { _ = "STUB: not implemented"; return nil }

func (s *receiver) EncodeArray(f func(Encoder) error) error { _ = "STUB: not implemented"; return nil }

func (s *receiver) EncodeField(field string, f func(Encoder) error) error {
	_ = "STUB: not implemented"
	return nil
}

type arrayReceiver struct {
	set   bool
	items []string
}

func (e *arrayReceiver) EncodeValue(v string) error { _ = "STUB: not implemented"; return nil }

func (e *arrayReceiver) EncodeArray(_ func(Encoder) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *arrayReceiver) EncodeField(_ string, _ func(Encoder) error) error {
	_ = "STUB: not implemented"
	return nil
}

type valueReceiver struct {
	set   bool
	value string
}

func (e *valueReceiver) EncodeValue(v string) error { _ = "STUB: not implemented"; return nil }

func (e *valueReceiver) EncodeArray(_ func(Encoder) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *valueReceiver) EncodeField(_ string, _ func(Encoder) error) error {
	_ = "STUB: not implemented"
	return nil
}
