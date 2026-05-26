package ir

type GenericVariant struct {
	Nullable bool
	Optional bool
}

func (v GenericVariant) NullableOptional() bool { _ = "STUB: not implemented"; return false }

func (v GenericVariant) OnlyOptional() bool { _ = "STUB: not implemented"; return false }

func (v GenericVariant) OnlyNullable() bool { _ = "STUB: not implemented"; return false }

func (v GenericVariant) Name() string { _ = "STUB: not implemented"; return "" }

func (v GenericVariant) Any() bool { _ = "STUB: not implemented"; return false }

// CanGeneric reports whether Type can be boxed to KindGeneric.
func (t Type) CanGeneric() bool { _ = "STUB: not implemented"; return false }
