package ir

import (
	"github.com/ogen-go/ogen/internal/bitset"
)

// SecurityKind defines security kind.
type SecurityKind string

const (
	// QuerySecurity is URL query security parameter. Matches "apiKey" type with "in" = "query".
	QuerySecurity SecurityKind = "query"
	// HeaderSecurity is HTTP header security parameter. Matches some "http" schemes and "apiKey" with "in" = "header".
	HeaderSecurity SecurityKind = "header"
	// CookieSecurity is HTTP cookie security parameter. Matches some "http" schemes and "apiKey" with "in" = "cookie".
	CookieSecurity SecurityKind = "cookie"
	// OAuth2Security is special type for OAuth2-based authentication. Matches "oauth2" and "openIdConnect".
	OAuth2Security SecurityKind = "oauth2"
)

// IsQuery whether s is QuerySecurity.
func (s SecurityKind) IsQuery() bool { _ = "STUB: not implemented"; return false }

// IsHeader whether s is HeaderSecurity.
func (s SecurityKind) IsHeader() bool { _ = "STUB: not implemented"; return false }

// IsCookie whether s is CookieSecurity.
func (s SecurityKind) IsCookie() bool { _ = "STUB: not implemented"; return false }

// IsOAuth2 whether s is OAuth2Security.
func (s SecurityKind) IsOAuth2() bool { _ = "STUB: not implemented"; return false }

// SecurityFormat defines security parameter format.
type SecurityFormat string

const (
	// APIKeySecurityFormat is plain value format.
	APIKeySecurityFormat SecurityFormat = "apiKey"
	// BearerSecurityFormat is Bearer authentication (RFC 6750) format.
	//
	// Unsupported yet.
	BearerSecurityFormat SecurityFormat = "bearer"
	// BasicHTTPSecurityFormat is Basic HTTP authentication (RFC 7617) format.
	BasicHTTPSecurityFormat SecurityFormat = "basic"
	// DigestHTTPSecurityFormat is Digest HTTP authentication (RFC 7616) format.
	//
	// Unsupported yet.
	DigestHTTPSecurityFormat SecurityFormat = "digest"

	// Oauth2SecurityFormat is Oauth2 security format.
	Oauth2SecurityFormat SecurityFormat = "oauth2"

	// CustomSecurityFormat is a user-defined security format.
	CustomSecurityFormat = "x-ogen-custom-security"
)

// IsAPIKeySecurity whether s is APIKeySecurityFormat.
func (s SecurityFormat) IsAPIKeySecurity() bool { _ = "STUB: not implemented"; return false }

// IsBearerSecurity whether s is BearerSecurityFormat.
func (s SecurityFormat) IsBearerSecurity() bool { _ = "STUB: not implemented"; return false }

// IsBasicHTTPSecurity whether s is BasicHTTPSecurityFormat.
func (s SecurityFormat) IsBasicHTTPSecurity() bool { _ = "STUB: not implemented"; return false }

// IsDigestHTTPSecurity whether s is DigestHTTPSecurityFormat.
func (s SecurityFormat) IsDigestHTTPSecurity() bool { _ = "STUB: not implemented"; return false }

// IsOAuth2Security whether s is Oauth2SecurityFormat.
func (s SecurityFormat) IsOAuth2Security() bool { _ = "STUB: not implemented"; return false }

// IsCustomSecurity whether s is CustomSecurityFormat.
func (s SecurityFormat) IsCustomSecurity() bool { _ = "STUB: not implemented"; return false }

type Security struct {
	Kind          SecurityKind
	Format        SecurityFormat
	ParameterName string
	Description   string
	Type          *Type
	Scopes        map[string][]string
}

func (s *Security) GoDoc() []string { _ = "STUB: not implemented"; return nil }

type SecurityRequirements struct {
	Securities   []*Security
	Requirements []bitset.Bitset
}

// BitArrayLen returns the length for bitset's underlying array.
func (s SecurityRequirements) BitArrayLen() (r int) { _ = "STUB: not implemented"; return 0 }
