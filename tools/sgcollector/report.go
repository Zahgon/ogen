package main

import (
	"context"
	"crypto/sha256"
)

type Report struct {
	File           FileMatch
	Error          string            `json:",omitempty"`
	NotImplemented []string          `json:",omitempty"`
	Hash           [sha256.Size]byte `json:"-"`
}

type Reporter struct {
	stage   Stage
	ch      chan Report
	counter int
}

func (r *Reporter) run(ctx context.Context, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reporter) close() { _ = "STUB: not implemented"; return }

type Reporters struct {
	reporters [last]*Reporter
}

func (r *Reporters) init(buf int) {
	for i := range r.reporters {
		r.reporters[i] = &Reporter{
			stage:   Stage(i),
			ch:      make(chan Report, buf),
			counter: 0,
		}
	}
}

func (r *Reporters) close() { _ = "STUB: not implemented"; return }

func (r *Reporters) run(ctx context.Context, clean bool, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reporters) report(ctx context.Context, stage Stage, report Report) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reporters) writeStats(output string, total int) error {
	_ = "STUB: not implemented"
	return nil
}
