package location

// File is a source file.
type File struct {
	// Name is the file name.
	Name string
	// Source is the path or URL of the file.
	Source string
	// Lines stores newline offsets.
	Lines Lines
}

// HumanName returns human-friendly name for this File.
func (f File) HumanName() string { _ = "STUB: not implemented"; return "" }

// IsZero returns true if file has zero value.
func (f File) IsZero() bool { _ = "STUB: not implemented"; return false }

// File is not useful if lines is empty.

// NewFile creates a new File.
//
// Do not modify the data after calling this function, Lines will point to it.
func NewFile(name, source string, data []byte) File { _ = "STUB: not implemented"; return *new(File) }
