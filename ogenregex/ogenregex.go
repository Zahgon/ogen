// Package ogenregex provides an interface to the regex engine.
//
// JSON Schema specification prefers to use ECMA 262 regular expressions. However, Go's
// regex engine is based on RE2, which is a different engine. Also, Go's regex engine
// does not support lookbehind assertions, to ensure linear time matching.
//
// This package provides unified interface to both engines. Go's regex engine is used
// by default, but if the regex is not supported, the dlclark/regexp2 would be used.
package ogenregex

import (
	"regexp"

	"github.com/dlclark/regexp2"
)

var _ = []Regexp{
	goRegexp{},
	regexp2Regexp{},
}

type goRegexp struct {
	orig string
	exp  *regexp.Regexp
}

func (r goRegexp) Match(s []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (r goRegexp) MatchString(s string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (r goRegexp) String() string { _ = "STUB: not implemented"; return "" }

type regexp2Regexp struct {
	exp *regexp2.Regexp
}

func (r regexp2Regexp) Match(s []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (r regexp2Regexp) MatchString(s string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r regexp2Regexp) String() string { _ = "STUB: not implemented"; return "" }

// Regexp is a regular expression interface.
type Regexp interface {
	Match(s []byte) (bool, error)
	MatchString(s string) (bool, error)
	String() string
}

// Compile compiles a regular expression.
//
// NOTE: this function may compile the same expression multiple times and can
// be slow. Compile the expression once and reuse it.
func Compile(exp string) (Regexp, error) { _ = "STUB: not implemented"; return *new(Regexp), nil }

// FIXME(tdakkota): Default timeout is "forever", which may lead to DoS.
// 	Probably, we should make this configurable.

// MustCompile compiles a regular expression and panics on error.
//
// NOTE: this function may compile the same expression multiple times and can
// be slow. Compile the expression once and reuse it.
func MustCompile(exp string) Regexp { _ = "STUB: not implemented"; return *new(Regexp) }
