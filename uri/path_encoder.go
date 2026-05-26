package uri

type PathStyle string

const (
	PathStyleSimple PathStyle = "simple"
	PathStyleLabel  PathStyle = "label"
	PathStyleMatrix PathStyle = "matrix"
)

func (s PathStyle) String() string { _ = "STUB: not implemented"; return "" }

type PathEncoder struct {
	param   string    // immutable
	style   PathStyle // immutable
	explode bool      // immutable
	*receiver
}

type PathEncoderConfig struct {
	Param   string
	Style   PathStyle
	Explode bool
}

func NewPathEncoder(cfg PathEncoderConfig) *PathEncoder { _ = "STUB: not implemented"; return nil }

func (e *PathEncoder) checkParam() error { _ = "STUB: not implemented"; return nil }

func (e *PathEncoder) Result() (r string, _ error) { _ = "STUB: not implemented"; return "", nil }

func (e *PathEncoder) value() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (e *PathEncoder) array() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (e *PathEncoder) object() (string, error) { _ = "STUB: not implemented"; return "", nil }
