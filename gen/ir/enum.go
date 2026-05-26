package ir

type EnumVariant struct {
	Name  string
	Value any
}

func (v *EnumVariant) ValueGo() string { _ = "STUB: not implemented"; return "" }
