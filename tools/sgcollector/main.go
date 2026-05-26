package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

func run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// get up-to-date statistics

// Wait until all writers stopped.

// Close readers.

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := run(ctx); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
