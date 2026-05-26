package main

// noVerboseError is an error that doesn't print the stack trace in zap.
type noVerboseError struct {
	err error
}

func (n noVerboseError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (n noVerboseError) Error() string { _ = "STUB: not implemented"; return "" }

// GenerateError reports that generation failed.
type GenerateError struct {
	stage   Stage
	notImpl []string
	err     error
}

func (p *GenerateError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (p *GenerateError) Error() string { _ = "STUB: not implemented"; return "" }
