package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"go.uber.org/zap"
)

// Run calls f and exists with code 2 on error.
func Run(f func(ctx context.Context, lg *zap.Logger) error) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	lg, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	if err := f(ctx, lg); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed: %+v\n", err)
		os.Exit(2)
	}
}
