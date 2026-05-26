package jsonpointer

import (
	"net/url"

	"github.com/ogen-go/ogen/location"
)

// RefKey is JSON Reference key.
type RefKey struct {
	// Loc is an URL of JSON document.
	Loc string
	// Ptr is JSON Pointer.
	Ptr string
}

// String returns string representation of reference.
func (r RefKey) String() string { _ = "STUB: not implemented"; return "" }

// IsZero returns true if RefKey is zero.
func (r RefKey) IsZero() bool { _ = "STUB: not implemented"; return false }

// FromURL sets RefKey from URL.
func (r *RefKey) FromURL(u *url.URL) {
	_ = "STUB: not implemented"

	// Make copy.
	return
}

type locstackItem struct {
	loc  *url.URL
	file location.File
}

// ResolveCtx is JSON Reference resolve context.
type ResolveCtx struct {
	// Location stack. Used for context-depending resolving.
	//
	// For resolve trace like
	//
	// 	"#/components/schemas/Schema" ->
	// 	"https://example.com/schema#Schema" ->
	//	"#/definitions/SchemaProperty"
	//
	// "#/definitions/SchemaProperty" should be resolved against "https://example.com/schema".
	locstack []locstackItem
	// root is root location.
	root *url.URL
	// Store references to detect infinite recursive references.
	refs       map[RefKey]struct{}
	depthLimit int
}

// DefaultDepthLimit is default depth limit for ResolveCtx.
const DefaultDepthLimit = 1000

// DummyURL is dummy URL for testing purposes.
func DummyURL() *url.URL { _ = "STUB: not implemented"; return nil }

// NewResolveCtx creates new ResolveCtx.
func NewResolveCtx(root *url.URL, depthLimit int) *ResolveCtx {
	_ = "STUB: not implemented"
	return nil
}

func (r ResolveCtx) last() (last locstackItem, ok bool) {
	_ = "STUB: not implemented"
	return *new(locstackItem), false
}

// Key creates new reference key.
func (r *ResolveCtx) Key(ref string) (key RefKey, _ error) {
	_ = "STUB: not implemented"
	return *new(RefKey), nil
}

// AddKey adds reference key to context.
func (r *ResolveCtx) AddKey(key RefKey, file location.File) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete removes reference from context.
func (r *ResolveCtx) Delete(key RefKey) { _ = "STUB: not implemented"; return }

// IsRoot returns true if location stack is empty.
func (r *ResolveCtx) IsRoot(key RefKey) bool { _ = "STUB: not implemented"; return false }

// File returns last file from stack.
func (r *ResolveCtx) File() (f location.File) {
	_ = "STUB: not implemented"
	return *new(location.File)
}
