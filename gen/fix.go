package gen

import (
	"github.com/ogen-go/ogen/gen/ir"
)

// Example:
//
//	responses:
//		200:
//		  contents:
//		    application/json:
//		      ref: #/components/schemas/Foo
//		202:
//		  contents:
//		    application/json:
//		      ref: #/components/schemas/Foo
//
// This response refers to the same schema for different
// status codes, and it will cause a collision:
//
//	func encodeResponse(resp FooResponse) {
//	    switch resp.(type) {
//		case *Foo:
//	    case *Foo:
//	    }
//	}
//
// To prevent collision we wrap referenced schema with aliases
// and use them instead.
//
//	type FooResponseOK Foo
//	func(*FooResponseOK) FooResponse() {}
//
//	type FooResponseAccepted Foo
//	func(*FooResponseAccepted) FooResponse() {}
//
// Referring to the same schema in different content types
// also can cause a collision and it will be fixed in the same way.
func fixEqualResponses(ctx *genctx, op *ir.Operation) error { _ = "STUB: not implemented"; return nil }

// We can modify contents of operation response.
// To prevent changes affecting to other operations
// (in case of referenced responses), we copy the response.

// Add `Content-Type` to response name only if needed.

// TODO: Fix duplicates.
// g.saveType(alias)

func cloneResponse(r *ir.Responses) *ir.Responses { _ = "STUB: not implemented"; return nil }

func fixEqualRequests(ctx *genctx, op *ir.Operation) error { _ = "STUB: not implemented"; return nil }

// We can modify request contents.
// To prevent changes affecting to other operations
// (in case of referenced requestBodies), we copy requestBody.

// TODO: Fix duplicates.
// g.saveType(alias)

func cloneRequest(r *ir.Request) *ir.Request { _ = "STUB: not implemented"; return nil }
