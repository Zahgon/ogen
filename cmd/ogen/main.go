// Binary ogen generates go source code from OAS.
package main

import (
	"fmt"
	"io"
	"os"

	"go.uber.org/zap"

	"github.com/ogen-go/ogen/gen"
)

func cleanDir(targetDir string, files []os.DirEntry) (rerr error) {
	_ = "STUB: not implemented"
	return nil
}

// Do not return error if file does not exist.
//#nosec G703

// Do not stop on first error, try to remove all files.

func generate(data []byte, packageName, targetDir string, clean bool, opts gen.Options) error {
	_ = "STUB: not implemented"
	return nil
}

// For pretty error message, we need to pass location.File.

// Clean target dir only after flag parsing, spec parsing and IR building.

//#nosec G703

// FIXME(tdakkota): write source uses imports.Process which also uses go/format.
// 	So, there is no reason to format source twice or provide a flag to disable formatting.

func handleGenerateError(w io.Writer, color bool, err error) (r bool) {
	_ = "STUB: not implemented"

	// Add trailing newline to the error message if error is handled.
	return false
}

//#nosec 6705

func handleNotImplementedError(err error) (msg, feature string, _ bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

//#nosec G705

//#nosec G705

//#nosec G705

// 10

// Sort by number of 'also used' types.
//
// It is likely to be properties to be fixed.

//#nosec G705

//#nosec G705

//#nosec G705

func loadConfig(cfgPath string, log *zap.Logger) (opts gen.Options, _ error) {
	_ = "STUB: not implemented"
	return *new(gen.Options), nil
}

//#nosec G703

func run() error { _ = "STUB: not implemented"; return nil }

//#nosec G705

// Config flag.

// Generator options.

// Parser options.

// Logging options.

// Profile options.

// Version option.

//#nosec G703

//#nosec G703

// Apply CLI flags that override config

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
