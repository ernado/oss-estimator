package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/go-faster/errors"
	"go.uber.org/zap"

	"estimator/internal/app"
	"estimator/internal/gh"
)

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		var arg struct {
			User string
		}
		flag.StringVar(&arg.User, "user", "ernado", "user name")
		flag.Parse()

		c := gh.Client(ctx)
		u, _, err := c.Users.Get(ctx, "go-faster")
		if err != nil {
			return errors.Wrap(err, "get user")
		}

		fmt.Println("ID:", u.GetID())
		company := u.GetCompany()
		if company == "" {
			company = "Independent"
		}
		fmt.Println("Company:", company)
		return nil
	})
}
