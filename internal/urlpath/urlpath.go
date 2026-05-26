// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code copied from
// https://github.com/golang/go/blob/2ebe77a2fda1ee9ff6fd9a3e08933ad1ebaea039/src/cmd/go/internal/web/url.go
// TODO(go.dev/issue/32456): if accepted, use the new API.

// Package urlpath provides utilities for converting between file URLs and file paths.
package urlpath

import (
	"errors"
	"net/url"
)

var errNotAbsolute = errors.New("path is not absolute")

// URLToFilePath converts a file-scheme url to a file path.
func URLToFilePath(u *url.URL) (string, error) { _ = "STUB: not implemented"; return "", nil }

// URLFromFilePath converts the given absolute path to a URL.
func URLFromFilePath(path string) (*url.URL, error) { _ = "STUB: not implemented"; return nil, nil }

// If path has a Windows volume name, convert the volume to a host and prefix
// per https://blogs.msdn.microsoft.com/ie/2006/12/06/file-uris-in-windows/.

// A degenerate case.
// \\host.example.com (without a share name)
// becomes
// file://host.example.com/

// \\host.example.com\Share\path\to\file
// becomes
// file://host.example.com/Share/path/to/file

// C:\path\to\file
// becomes
// file:///C:/path/to/file

// /path/to/file
// becomes
// file:///path/to/file

func convertFileURLPath(host, path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func convertFileURLPathWindows(host, path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// We interpret Windows file URLs per the description in
// https://blogs.msdn.microsoft.com/ie/2006/12/06/file-uris-in-windows/.

// The host part of a file URL (if any) is the UNC volume name,
// but RFC 8089 reserves the authority "localhost" for the local machine.

// A common "legacy" format omits the leading slash before a drive letter,
// encoding the drive letter as the host instead of part of the path.
// (See https://blogs.msdn.microsoft.com/freeassociations/2005/05/19/the-bizarre-and-unhappy-story-of-file-urls/.)
// We do not support that format, but we should at least emit a more
// helpful error message for it.

// If host is empty, path must contain an initial slash followed by a
// drive letter and path. Remove the slash and verify that the path is valid.
