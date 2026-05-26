package otelogen

import (
	"go.opentelemetry.io/otel/attribute"
)

const (
	// OperationIDKey by OpenAPI specification.
	OperationIDKey = attribute.Key("oas.operation")
	// WebhookNameKey by OpenAPI specification.
	WebhookNameKey = attribute.Key("oas.webhook.name")
)

// OperationID attribute.
func OperationID(v string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

// WebhookName attribute.
func WebhookName(v string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
