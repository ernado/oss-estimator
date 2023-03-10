package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Run calls f and exists with code 2 on error.
func Run(f func(ctx context.Context, lg *zap.Logger) error) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	lg := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg),
		zapcore.AddSync(colorable.NewColorableStdout()),
		zapcore.DebugLevel,
	))

	if err := f(ctx, lg); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed: %+v\n", err)
		os.Exit(2)
	}
}
