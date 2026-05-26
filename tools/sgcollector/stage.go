package main

type Stage int

func (s Stage) OnlyCounter() bool { _ = "STUB: not implemented"; return false }

const (
	InvalidYAML Stage = iota
	InvalidJSON
	Unmarshal
	Parse
	BuildIR
	BuildRouter
	Template
	Format
	NotImplemented
	Good
	Crash
	last
)

func (s Stage) String() string { _ = "STUB: not implemented"; return "" }

// #nosec G602
