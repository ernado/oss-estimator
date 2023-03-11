package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ClickHouse/ch-go"
	"github.com/ClickHouse/ch-go/proto"
	"github.com/go-faster/errors"
	"go.uber.org/zap"

	"estimator/internal/app"
)

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		c, err := ch.Dial(ctx, ch.Options{
			Database:    "estimator",
			Logger:      lg,
			Compression: ch.CompressionLZ4,
		})
		if err != nil {
			return errors.Wrap(err, "dial")
		}
		defer func() { _ = c.Close() }()

		var (
			colEv      proto.ColEnum8
			colRepoID  proto.ColInt64
			colActorID proto.ColInt64
			colTime    proto.ColDateTime

			total int
		)
		start := time.Now()
		if err := c.Do(ctx, ch.Query{
			Body: "SELECT * FROM events LIMIT 10000000",
			OnResult: func(ctx context.Context, block proto.Block) error {
				lg.Info("Block", zap.Int("rows", block.Rows))
				total += block.Rows
				return nil
			},
			Result: proto.Results{
				{Name: "event", Data: &colEv},
				{Name: "repo", Data: &colRepoID},
				{Name: "actor", Data: &colActorID},
				{Name: "time", Data: &colTime},
			},
		}); err != nil {
			return errors.Wrap(err, "query")
		}

		fmt.Println("Done, rps:", int64(float64(total)/time.Since(start).Seconds()))

		return nil
	})
}
