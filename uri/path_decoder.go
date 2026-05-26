package uri

type PathDecoder struct {
	cur *cursor

	param   string
	style   PathStyle
	explode bool
}

type PathDecoderConfig struct {
	Param   string // Parameter name
	Value   string // chi.URLParam(r, "paramName")
	Style   PathStyle
	Explode bool
}

func NewPathDecoder(cfg PathDecoderConfig) *PathDecoder { _ = "STUB: not implemented"; return nil }

func (d *PathDecoder) DecodeValue() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (d *PathDecoder) DecodeArray(f func(d Decoder) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *PathDecoder) DecodeFields(f func(name string, d Decoder) error) error {
	_ = "STUB: not implemented"
	return nil
}

func parseArray(cur *cursor, delim byte, f func(d Decoder) error) error {
	_ = "STUB: not implemented"
	return nil
}
