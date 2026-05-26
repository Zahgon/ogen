package gen

import (
	"github.com/go-faster/yaml"
)

// Feature is an ogen feature.
type Feature struct {
	Name        string
	Description string
}

// FeatureOptions is features Options.
type FeatureOptions struct {
	Enable     FeatureSet `json:"enable" yaml:"enable"`
	Disable    FeatureSet `json:"disable" yaml:"disable"`
	DisableAll bool       `json:"disable_all" yaml:"disable_all"`
}

// Build returns final set.
func (cfg *FeatureOptions) Build() (set FeatureSet, _ error) {
	_ = "STUB: not implemented"
	return *new(FeatureSet), nil
}

// FeatureSet is set of [Feature] names.
type FeatureSet map[string]struct{}

// Enable adds a feature to set.
func (s *FeatureSet) Enable(name string) error { _ = "STUB: not implemented"; return nil }

// Disable removes a feature from set.
func (s *FeatureSet) Disable(name string) {
	_ = "STUB: not implemented"

	// Has whether if set has given feature.
	return
}

func (s FeatureSet) Has(feature Feature) bool { _ = "STUB: not implemented"; return false }

// UnmarshalYAML implements [yaml.Unmarshaler].
func (s *FeatureSet) UnmarshalYAML(n *yaml.Node) error { _ = "STUB: not implemented"; return nil }

var (
	PathsClient = Feature{
		"paths/client",
		`Enables paths client generation`,
	}
	PathsServer = Feature{
		"paths/server",
		`Enables paths server generation`,
	}
	WebhooksClient = Feature{
		"webhooks/client",
		`Enables webhooks client generation`,
	}
	WebhooksServer = Feature{
		"webhooks/server",
		`Enables webhooks server generation`,
	}
	ClientSecurityReentrant = Feature{
		"client/security/reentrant",
		`Enables client usage in security source implementations`,
	}
	ClientRequestOptions = Feature{
		"client/request/options",
		`Enables function options for client requests`,
	}
	ClientRequestValidation = Feature{
		"client/request/validation",
		`Enables validation of client requests`,
	}
	ClientEditors = Feature{
		"client/editors",
		`Enables editors function options for client`,
	}
	ServerResponseValidation = Feature{
		"server/response/validation",
		`Enables validation of server responses`,
	}
	OgenOtel = Feature{
		"ogen/otel",
		`Enables OpenTelemetry integration`,
	}
	OgenUnimplemented = Feature{
		"ogen/unimplemented",
		`Enables stub Handler generation`,
	}
	DebugExampleTests = Feature{
		"debug/example_tests",
		`Enables example tests generation`,
	}
)

// DefaultFeatures defines default ogen features.
var DefaultFeatures = []Feature{
	PathsClient,
	PathsServer,
	WebhooksClient,
	WebhooksServer,
	OgenOtel,
	OgenUnimplemented,
}

// AllFeatures contains all ogen features.
var AllFeatures = []Feature{
	PathsClient,
	PathsServer,
	WebhooksClient,
	WebhooksServer,
	ClientSecurityReentrant,
	ClientRequestOptions,
	ClientRequestValidation,
	ClientEditors,
	ServerResponseValidation,
	OgenOtel,
	OgenUnimplemented,
	DebugExampleTests,
}
