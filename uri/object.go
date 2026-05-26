package uri

type Field struct {
	Name  string
	Value string
}

func encodeObject(kvSep, fieldSep byte, fields []Field) string {
	_ = "STUB: not implemented"
	return ""
}

// Preallocate the buffer.

// If there are less than 2 fields, we don't need to add the field separator.

func decodeObject(cur *cursor, kvSep, fieldSep byte, f func(field, value string) error) error {
	_ = "STUB: not implemented"
	return nil
}
