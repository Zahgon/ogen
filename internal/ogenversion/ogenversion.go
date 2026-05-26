// Package ogenversion provides the version of the ogen tool.
package ogenversion

import (
	"runtime/debug"
	"sync"
	"time"
)

var getOnce struct {
	info Info
	ok   bool
	once sync.Once
}

func getOgenVersion(m *debug.Module) (string, bool) { _ = "STUB: not implemented"; return "", false }

func getInfo() (Info, bool) { _ = "STUB: not implemented"; return *new(Info), false }

// ogen is the main module, so we can use buildvcs data.

// Info is the ogen build information.
type Info struct {
	// Version is the version of the ogen tool.
	Version string
	// GoVersion is the version of the Go that produced this binary.
	GoVersion string

	// Commit is the current commit hash.
	Commit string
	// Time is the time of the build.
	Time time.Time
}

// GetInfo returns the ogen build information.
func GetInfo() (Info, bool) {
	_ = "STUB: not implemented"

	// String returns string representation of the build information.
	return *new(Info), false
}

func (i Info) String() string { _ = "STUB: not implemented"; return "" }
