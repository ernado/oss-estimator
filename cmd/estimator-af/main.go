package main

import (
	"context"
	"os"

	"github.com/go-faster/jx"
	"go.uber.org/zap"

	"estimator/internal/af"
	"estimator/internal/app"
)

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		var e jx.Encoder
		return af.Parse(os.Stdin, func(u af.User) error {
			e.Reset()
			e.Obj(func(e *jx.Encoder) {
				e.Field("name", func(e *jx.Encoder) {
					e.Str(u.Name)
				})
				e.Field("emails", func(e *jx.Encoder) {
					e.Arr(func(e *jx.Encoder) {
						for _, email := range u.Emails {
							e.Str(email)
						}
					})
				})
				e.Field("companies", func(e *jx.Encoder) {
					e.Arr(func(e *jx.Encoder) {
						for _, company := range u.Companies {
							e.Obj(func(e *jx.Encoder) {
								e.Field("name", func(e *jx.Encoder) {
									e.Str(company.Name)
								})
								if !company.From.IsZero() {
									e.Field("from", func(e *jx.Encoder) {
										e.Str(company.From.Format("2006-01-02"))
									})
								}
								if !company.Until.IsZero() {
									e.Field("until", func(e *jx.Encoder) {
										e.Str(company.Until.Format("2006-01-02"))
									})
								}
							})
						}
					})
				})
			})

			if _, err := os.Stdout.Write(e.Bytes()); err != nil {
				return err
			}
			if _, err := os.Stdout.Write([]byte{'\n'}); err != nil {
				return err
			}
			return nil
		})
	})
}
