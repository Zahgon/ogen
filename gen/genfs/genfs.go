// Package genfs contains gen.FileSystem implementations.
package genfs

// FormattedSource is gen.FileSystem implementation that format and writes Go sources.
type FormattedSource struct {
	Format bool
	Root   string
}

// WriteFile implements gen.FileSystem.
func (t FormattedSource) WriteFile(name string, content []byte) error {
	_ = "STUB: not implemented"
	return nil
}
