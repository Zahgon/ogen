package uri

func ishex(c byte) bool { _ = "STUB: not implemented"; return false }

func unhex(c byte) byte { _ = "STUB: not implemented"; return 0 }

func asciiToUpper(c byte) byte { _ = "STUB: not implemented"; return 0 }

func asciiIsLowercase(c byte) bool { _ = "STUB: not implemented"; return false }

// Return true if the specified character should be escaped when
// appearing in a URL path string, according to RFC 3986.
func shouldEscapePath(c byte) bool {
	_ = "STUB: not implemented"
	// §2.3 Unreserved characters (alpha)
	return false
}

// 0-9

// §2.3 Unreserved characters (mark)

// Everything else must be escaped.

// NormalizeEscapedPath normalizes escaped path.
//
// All percent-encoded characters are upper-cased. If s contains unnecessarily escaped
// characters, they are unescaped.
//
// If s contains invalid escape sequence, it returns empty string and false.
func NormalizeEscapedPath(s string) (string, bool) {
	_ = "STUB: not implemented"
	// Search % with lower case octets.
	return "", false
}

// Invalid escape sequence.

// Unescape character.

// Unescape character.

// Unescape character.
