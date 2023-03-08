package main

import (
	"bufio"
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	lru "github.com/hashicorp/golang-lru/v2"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"estimator/internal/app"
	"estimator/internal/archive"
	"estimator/internal/entry"
)

type SimplifiedEntry struct {
	Event   byte
	Repo    int64
	ActorID int64
	Actor   []byte
	Time    int64
}

func (s SimplifiedEntry) Encode(e *jx.Encoder) {
	e.Obj(func(e *jx.Encoder) {
		e.Field("e", func(e *jx.Encoder) {
			e.UInt8(s.Event)
		})
		e.Field("r", func(e *jx.Encoder) {
			e.Int64(s.Repo)
		})
		e.Field("aid", func(e *jx.Encoder) {
			e.Int64(s.ActorID)
		})
		e.Field("a", func(e *jx.Encoder) {
			e.ByteStr(s.Actor)
		})
		e.Field("t", func(e *jx.Encoder) {
			e.Int64(s.Time)
		})
	})
}

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		u := archive.GetURL(time.Now().AddDate(0, 0, -2))

		g, ctx := errgroup.WithContext(ctx)

		users, err := bbolt.Open(filepath.Join("_work", "users.bbolt"), 0666, &bbolt.Options{NoSync: true})
		if err != nil {
			return errors.Wrap(err, "db open")
		}
		defer func() {
			lg.Info("Closing")
			_ = users.Sync()
			_ = users.Close()
		}()

		cache, err := lru.New[int64, string](5000)
		if err != nil {
			return errors.Wrap(err, "cache")
		}

		wget := exec.CommandContext(ctx, "wget",
			"-nv", "-O", "-", u,
		)
		wget.Stderr = os.Stderr
		dl, err := wget.StdoutPipe()
		if err != nil {
			return err
		}

		cmd := exec.CommandContext(ctx, "zstd", "-d")
		cmd.Stdin = dl
		cmd.Stderr = os.Stderr
		out, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}

		done := make(chan struct{})

		g.Go(func() error {
			if err := cmd.Start(); err != nil {
				return errors.Wrap(err, "zstd")
			}
			select {
			case <-ctx.Done():
				if err := cmd.Wait(); err != nil {
					return errors.Wrap(err, "zstd")
				}
			case <-done:
			}
			return nil
		})
		g.Go(func() error {
			if err := wget.Start(); err != nil {
				return errors.Wrap(err, "wget")
			}
			select {
			case <-ctx.Done():
				if err := wget.Wait(); err != nil {
					return errors.Wrap(err, "wget")
				}
			case <-done:
			}
			return nil
		})
		g.Go(func() error {
			defer close(done)
			s := bufio.NewScanner(out)
			var (
				d  jx.Decoder
				e  jx.Encoder
				ev entry.Event
			)
			buf := make([]byte, 0, 1024*1024)
			s.Buffer(buf, len(buf))

			nl := []byte{'\n'}
			bucket := []byte("id-to-actor")
			bucketInverse := []byte("actor-to-id")
			var idEncoder jx.Encoder

			for s.Scan() {
				e.Reset()
				ev.Reset()

				data := s.Bytes()
				d.ResetBytes(data)

				if err := ev.Decode(&d); err != nil {
					lg.Error("decode", zap.Error(err))
					continue
				}
				if !ev.Interesting() {
					continue
				}

				if found, _ := cache.ContainsOrAdd(ev.ActorID, string(ev.Actor)); !found {
					if err := users.Update(func(tx *bbolt.Tx) error {
						idEncoder.Reset()
						idEncoder.Int64(ev.ActorID)
						id := idEncoder.Bytes()
						b, err := tx.CreateBucketIfNotExists(bucket)
						if err != nil {
							return err
						}
						if err := b.Put(id, ev.Actor); err != nil {
							return err
						}
						bi, err := tx.CreateBucketIfNotExists(bucketInverse)
						if err != nil {
							return err
						}
						if err := bi.Put(ev.Actor, id); err != nil {
							return err
						}
						return nil
					}); err != nil {
						return errors.Wrap(err, "db")
					}
				}

				se := SimplifiedEntry{
					Event:   byte(ev.Type),
					Repo:    ev.RepoID,
					ActorID: ev.ActorID,
					Actor:   ev.Actor,
					Time:    ev.Time.Unix(),
				}
				se.Encode(&e)
				if _, err := os.Stdout.Write(e.Bytes()); err != nil {
					return errors.Wrap(err, "write")
				}
				if _, err := os.Stdout.Write(nl); err != nil {
					return errors.Wrap(err, "write")
				}
			}
			if err := s.Err(); err != nil {
				return errors.Wrap(err, "scan")
			}
			return nil
		})

		return g.Wait()
	})
}
