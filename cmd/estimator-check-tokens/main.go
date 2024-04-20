package main

import (
	"context"

	"go.uber.org/zap"

	"estimator/internal/app"
	"estimator/internal/gh"
)

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		return gh.Check()
	})
}
