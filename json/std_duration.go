package json

import "time"

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// formatDuration is port of time.Duration String method to format into byte slice.
//
// It returns the offset from the end of passed buf.
func formatDuration(buf *[32]byte, d time.Duration) (w int) {
	_ = "STUB: not implemented"
	// Largest time is 2540400h10m10.000000000s
	return 0
}

// Special case: if duration is smaller than a second,
// use smaller units, like 1.2ms

// Note that formatted text must be in the right part of the buffer.
//
// So offset is len(buf) - n.

// print nanoseconds

// print microseconds

// U+00B5 'µ' micro sign == 0xC2 0xB5
// Need room for two bytes.

// print milliseconds

// u is now integer seconds

// u is now integer minutes

// u is now integer hours
// Stop at hours because days can be different lengths.

// fmtFrac formats the fraction of v/10**prec (e.g., ".12345") into the
// tail of buf, omitting trailing zeros. It omits the decimal
// point too when the fraction is 0. It returns the index where the
// output bytes begin and the value v/10**prec.
func fmtFrac(buf []byte, v uint64, prec int) (nw int, nv uint64) {
	_ = "STUB: not implemented"
	// Omit trailing zeros up to and including decimal point.
	return 0, 0
}

// fmtInt formats v into the tail of buf.
// It returns the index where the output begins.
func fmtInt(buf []byte, v uint64) int { _ = "STUB: not implemented"; return 0 }
