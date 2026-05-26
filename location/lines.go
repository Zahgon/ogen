package location

// Lines is a sorted slice of newline offsets.
type Lines struct {
	data []byte
	// lines stores newline offsets.
	//
	// idx is the line number (counts from 0).
	lines []int
}

// IsZero returns true if lines has zero value.
func (l Lines) IsZero() bool { _ = "STUB: not implemented"; return false }

// Line returns offset range of the line.
//
// NOTE: the line number is 1-based. Returns (-1, -1) if the line is invalid.
func (l Lines) Line(n int) (start, end int) { _ = "STUB: not implemented"; return 0, 0 }

// Line 0 is invalid.

// Last line.

// Collect fills the given slice with the offset of newlines.
func (l *Lines) Collect(data []byte) { _ = "STUB: not implemented"; return }

// Remaining data to process.

// Absolute offset of the current line.
