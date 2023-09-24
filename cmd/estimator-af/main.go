package main

import (
	"context"
	"flag"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/go-github/v50/github"
	"go.uber.org/zap"

	"estimator/internal/af"
	"estimator/internal/app"
	"estimator/internal/archive"
	"estimator/internal/gh"
)

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		var arg struct {
			Dir string
			DB  string
		}
		flag.StringVar(&arg.Dir, "dir", "", "directory to read from")
		flag.StringVar(&arg.DB, "db", "", "path to bbolt user database")
		flag.Parse()

		if err := gh.Check(); err != nil {
			return errors.Wrap(err, "check tokens")
		}

		c := gh.Client()
		db, err := archive.NewUserCache(archive.UserCacheOptions{
			Path: arg.DB,
		})
		if err != nil {
			return errors.Wrap(err, "open user cache")
		}
		defer func() {
			_ = db.Close()
		}()

		entries, err := os.ReadDir(arg.Dir)
		if err != nil {
			return errors.Wrap(err, "read dir")
		}

		var e jx.Encoder
		lg.Info("Listing")
		for _, entry := range entries {
			n := entry.Name()
			if !strings.HasPrefix(n, "developers_affiliations") {
				continue
			}
			if !strings.HasSuffix(n, ".txt") {
				continue
			}
			if ctx.Err() != nil {
				return ctx.Err()
			}
			if err := func() error {
				var total int
				defer func() {
					lg.Info("Processed",
						zap.String("file", n),
						zap.Int("total", total),
					)
				}()

				f, err := os.Open(filepath.Join(arg.Dir, n))
				if err != nil {
					return errors.Wrap(err, "open file")
				}
				defer func() { _ = f.Close() }()
				lg.Info("Processing", zap.String("file", n))

				return af.Parse(f, func(u af.User) error {
					total++

					if ctx.Err() != nil {
						return ctx.Err()
					}
					name := []byte(u.Name)
					id, err := db.Get(name)
					if err != nil {
						return errors.Wrap(err, "get user")
					}

					if id == 0 {
					Request:
						usr, res, err := c.Users.Get(ctx, u.Name)
						if re, ok := err.(*github.RateLimitError); ok {
							d := time.Until(re.Rate.Reset.Time) + time.Second*10
							lg.Warn("Rate limit exceeded",
								zap.Duration("reset", d),
								zap.Int("limit", re.Rate.Limit),
								zap.Int("remaining", re.Rate.Remaining),
							)
							select {
							case <-ctx.Done():
								return ctx.Err()
							case <-time.After(d):
								goto Request
							}
						}
						if err != nil {
							if res == nil || (res.StatusCode != 404 && res.StatusCode != 403) {
								return errors.Wrap(err, "get user")
							}
						} else {
							id = usr.GetID()
							lg.Info("Saving user to cache",
								zap.Int64("id", id),
								zap.String("name", string(name)),
							)
							if err := db.Add(id, name); err != nil {
								return errors.Wrap(err, "add user")
							}
						}
					}
					e.Reset()
					e.Obj(func(e *jx.Encoder) {
						e.Field("name", func(e *jx.Encoder) {
							e.Str(u.Name)
						})
						if id != 0 {
							e.Field("id", func(e *jx.Encoder) {
								e.Int64(id)
							})
						}
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
			}(); err != nil {
				return errors.Wrap(err, "process entry")
			}
		}

		return nil
	})
}
