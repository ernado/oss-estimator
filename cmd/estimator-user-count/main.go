package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.uber.org/zap"

	"estimator/internal/app"
)

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		var count int

		e := jx.Decode(os.Stdin, 1024)
		if err := e.Arr(func(d *jx.Decoder) error {
			count++
			return d.Skip()
		}); err != nil {
			return errors.Wrap(err, "decode")
		}

		fmt.Println("Count:", count)
		return nil
	})
}
