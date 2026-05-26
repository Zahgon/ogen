package uri

type constval struct {
	v string
}

func (d constval) DecodeValue() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (d constval) DecodeArray(f func(Decoder) error) error { _ = "STUB: not implemented"; return nil }

func (d constval) DecodeFields(f func(string, Decoder) error) error {
	_ = "STUB: not implemented"
	return nil
}
