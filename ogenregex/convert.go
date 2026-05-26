package ogenregex

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/rangetable"
)

// Copied from dop251/goja, to avoid dependency.
//
// All rights belong to the original author.
//
// https://github.com/dop251/goja/blob/3b8a68ca89b4fa7086a4236695032e10a69b2472/parser/regexp.go#L58

const (
	whitespaceChars = " \f\n\r\t\v" +
		"\u00a0\u1680" +
		"\u2000\u2001\u2002\u2003\u2004\u2005\u2006\u2007\u2008\u2009\u200a" +
		"\u2028\u2029" +
		"\u202f\u205f" +
		"\u3000\ufeff"
	re2Dot = "[^\r\n\u2028\u2029]"
)

func digitValue(chr rune) int { _ = "STUB: not implemented"; return 0 }

// Larger than any legal digit value

var (
	unicodeRangeIDNeg      = rangetable.Merge(unicode.Pattern_Syntax, unicode.Pattern_White_Space)
	unicodeRangeIDStartPos = rangetable.Merge(unicode.Letter, unicode.Nl, unicode.Other_ID_Start)
	unicodeRangeIDContPos  = rangetable.Merge(
		unicodeRangeIDStartPos,
		unicode.Mn,
		unicode.Mc,
		unicode.Nd,
		unicode.Pc,
		unicode.Other_ID_Continue,
	)
)

func isIDPartUnicode(r rune) bool { _ = "STUB: not implemented"; return false }

func isIdentifierPart(chr rune) bool { _ = "STUB: not implemented"; return false }

// Convert converts a ECMA-262 regular expression to Go's regular expression.
//
// If the conversion is not possible, ("", false) is returned.
func Convert(pattern string) (string, bool) { _ = "STUB: not implemented"; return "", false }

type parser struct {
	str    string
	length int

	chr       rune // The current character
	chrOffset int  // The offset of current character
	offset    int  // The offset after current character (may be greater than 1)

	err error

	goRegexp   strings.Builder
	passOffset int
}

func (p *parser) ResultString() string { _ = "STUB: not implemented"; return "" }

func (p *parser) parse() error {
	_ = "STUB: not implemented"
	// Pull in the first character
	return nil
}

func (p *parser) read() { _ = "STUB: not implemented"; return }

// !ASCII

// EOF

func (p *parser) stopPassing() { _ = "STUB: not implemented"; return }

func (p *parser) write(data []byte) { _ = "STUB: not implemented"; return }

func (p *parser) writeByte(b byte) { _ = "STUB: not implemented"; return }

func (p *parser) writeString(s string) { _ = "STUB: not implemented"; return }

func (p *parser) scan() { _ = "STUB: not implemented"; return }

// (...)
func (p *parser) scanGroup() { _ = "STUB: not implemented"; return }

// A possibility of (?= or (?!

// [...]
func (p *parser) scanBracket() { _ = "STUB: not implemented"; return }

// [] -- Empty character class

// \...
func (p *parser) scanEscape(inClass bool) { _ = "STUB: not implemented"; return }

// Not a valid digit

// The number of characters read

// An invalid backreference

// This is slightly broken, because ECMAScript
// includes \v in \s, \S, while re2 does not

// $ is an identifier character, so we have to have
// a special case for it here

// A non-identifier character needs escaping

// Unescape the character for re2

// Otherwise, we're a \u.... or \x...

// Not a valid digit

// Not a valid digit

// Should never, ever get here...

func (p *parser) pass() { _ = "STUB: not implemented"; return }

func (p *parser) passString(start, end int) { _ = "STUB: not implemented"; return }

func (p *parser) error(fatal bool, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}
