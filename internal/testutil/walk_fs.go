// Package testutil contains helper functions for testing.
package testutil

import (
	"io/fs"
	"testing"
)

// WalkTestdata recursively walks through the root directory in given testdata FS
// and spawns a test for each file.
func WalkTestdata(t *testing.T, testdata fs.FS, root string, cb func(t *testing.T, file string, data []byte)) {
	_ = "STUB: not implemented"
	return
}
