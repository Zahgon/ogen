package genfs

// CheckFS is in-memory gen.FileSystem implementation for checking Go sources.
type CheckFS struct{}

// WriteFile implements gen.FileSystem.
func (n CheckFS) WriteFile(baseName string, source []byte) error {
	_ = "STUB: not implemented"
	return nil
}
