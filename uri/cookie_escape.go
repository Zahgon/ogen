package uri

const cookieEscaper = '%'

var cookieEscapeChars = [128]byte{
	'\x00': 1,
	'\x01': 1,
	'\x02': 1,
	'\x03': 1,
	'\x04': 1,
	'\x05': 1,
	'\x06': 1,
	'\a':   1,
	'\b':   1,
	'\t':   1,
	'\n':   1,
	'\v':   1,
	'\f':   1,
	'\r':   1,
	'\x0e': 1,
	'\x0f': 1,
	'\x10': 1,
	'\x11': 1,
	'\x12': 1,
	'\x13': 1,
	'\x14': 1,
	'\x15': 1,
	'\x16': 1,
	'\x17': 1,
	'\x18': 1,
	'\x19': 1,
	'\x1a': 1,
	'\x1b': 1,
	'\x1c': 1,
	'\x1d': 1,
	'\x1e': 1,
	'\x1f': 1,
	' ':    1,
	'"':    1,
	',':    1,
	';':    1,
	'\\':   1,
	'\x7f': 1,

	// Escape the escape character itself.
	cookieEscaper: 1,
}

const hex = "0123456789ABCDEF"

func escapeCookie(s string) string { _ = "STUB: not implemented"; return "" }

// No need to escape.

// Every escaped char is 2 bytes longer: percent sign and 2 hex digits minus existing byte.

func unescapeCookie(s string) (string, bool) { _ = "STUB: not implemented"; return "", false }

// No need to unescape.

// Every escaped char is 2 bytes longer: percent sign and 2 hex digits minus existing byte.
