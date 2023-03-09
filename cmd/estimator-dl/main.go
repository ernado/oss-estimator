package main

import (
	"bufio"
	"context"
	"encoding/binary"
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/ClickHouse/ch-go/proto"
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/klauspost/compress/zstd"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"estimator/internal/app"
	"estimator/internal/archive"
	"estimator/internal/entry"
)

type UserCache struct {
	db  *bbolt.DB
	lru *lru.Cache[int64, struct{}]
}

var (
	_bucket        = []byte("id-to-actor")
	_bucketInverse = []byte("actor-to-id")
)

type key [8]byte

func (u *UserCache) inDB(k key) (found bool, err error) {
	if err := u.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(_bucket)
		if b == nil {
			return nil
		}
		found = b.Get(k[:]) != nil
		return nil
	}); err != nil {
		return false, errors.Wrap(err, "view")
	}
	return found, nil
}

func (u *UserCache) key(id int64) key {
	var k [8]byte
	binary.BigEndian.PutUint64(k[:], uint64(id))
	return k
}

func (u *UserCache) put(k key, v []byte) error {
	// Set both index values.
	if err := u.db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(_bucket)
		if err != nil {
			return errors.Wrap(err, "create bucket")
		}
		if err := b.Put(k[:], v); err != nil {
			return errors.Wrap(err, "put")
		}
		bi, err := tx.CreateBucketIfNotExists(_bucketInverse)
		if err != nil {
			return errors.Wrap(err, "create inverse bucket")
		}
		if err := bi.Put(v, k[:]); err != nil {
			return errors.Wrap(err, "put")
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "db")
	}
	return nil
}

func (u *UserCache) Close() error {
	u.lru.Purge()
	if err := u.db.Sync(); err != nil {
		return errors.Wrap(err, "sync")
	}
	if err := u.db.Close(); err != nil {
		return errors.Wrap(err, "close")
	}
	return nil
}

func (u *UserCache) hasOrAdd(id int64) bool {
	found, _ := u.lru.ContainsOrAdd(id, struct{}{})
	return found
}

func (u *UserCache) Add(id int64, v []byte) error {
	if u.hasOrAdd(id) {
		return nil
	}
	k := u.key(id)
	found, err := u.inDB(k)
	if err != nil {
		return errors.Wrap(err, "db get")
	}
	if found {
		// In DB, was already wrote.
		return nil
	}
	if err := u.put(k, v); err != nil {
		return errors.Wrap(err, "put")
	}
	return nil
}

func NewUserCache(dir string, size int) (*UserCache, error) {
	db, err := bbolt.Open(filepath.Join(dir, "users.bbolt"), 0666, &bbolt.Options{
		NoSync:         true,
		NoFreelistSync: true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "db open")
	}
	cache, err := lru.New[int64, struct{}](size)
	if err != nil {
		return nil, errors.Wrap(err, "cache")
	}
	return &UserCache{
		db:  db,
		lru: cache,
	}, nil
}

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		var arg = struct {
			Jobs int
			Dir  string
		}{
			Jobs: 8,
			Dir:  filepath.Join("_work"),
		}
		flag.IntVar(&arg.Jobs, "j", arg.Jobs, "concurrent jobs")
		flag.StringVar(&arg.Dir, "dir", arg.Dir, "directory to store data")
		flag.Parse()

		ce, err := NewUserCache(arg.Dir, 250_000)
		if err != nil {
			return errors.Wrap(err, "cache")
		}
		defer func() {
			lg.Info("Closing")
			_ = ce.Close()
		}()

		u := archive.GetURL(time.Now().AddDate(0, 0, -2))

		// clickhouse local --structure "event Enum8('WatchEvent'=1, 'PushEvent'=2, 'IssuesEvent'=3, 'PullRequestEvent'=4), repo Int64, actor Int64, time DateTime" --input-format Native --interactive --file events.ch.native.zst
		outPathTarget := filepath.Join(arg.Dir, "events.ch.native.zst")
		outPathTmp := outPathTarget + ".tmp"
		outFile, err := os.Create(outPathTmp)
		if err != nil {
			return errors.Wrap(err, "open")
		}
		defer func() { _ = outFile.Close() }()

		outWrite, err := zstd.NewWriter(outFile,
			zstd.WithEncoderConcurrency(1),
			zstd.WithEncoderLevel(zstd.SpeedBetterCompression),
		)
		if err != nil {
			return errors.Wrap(err, "zstd")
		}

		g, ctx := errgroup.WithContext(ctx)

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

				colEv      proto.ColEnum8
				colRepoID  proto.ColInt64
				colActorID proto.ColInt64
				colTime    proto.ColDateTime

				outBuf proto.Buffer
			)

			buf := make([]byte, 0, 1024*1024)
			s.Buffer(buf, len(buf))
			input := []proto.InputColumn{
				{Name: "event", Data: proto.Wrap(&colEv, `'WatchEvent'=1, 'PushEvent'=2, 'IssuesEvent'=3, 'PullRequestEvent'=4`)},
				{Name: "repo", Data: &colRepoID},
				{Name: "actor", Data: &colActorID},
				{Name: "time", Data: &colTime},
			}
			write := func() error {
				b := proto.Block{Rows: colEv.Rows(), Columns: 4}
				if err := b.EncodeRawBlock(&outBuf, 54451, input); err != nil {
					return errors.Wrap(err, "encode")
				}
				if _, err := outWrite.Write(outBuf.Buf); err != nil {
					return errors.Wrap(err, "write")
				}
				outBuf.Reset()
				proto.Reset(
					&colEv,
					&colRepoID,
					&colActorID,
					&colTime,
				)
				return nil
			}

			for s.Scan() {
				if ctx.Err() != nil {
					return ctx.Err()
				}

				e.Reset()
				ev.Reset()

				if colEv.Rows() > 10_000 {
					if err := write(); err != nil {
						return errors.Wrap(err, "write")
					}
				}

				data := s.Bytes()
				d.ResetBytes(data)

				if err := ev.Decode(&d); err != nil {
					lg.Error("decode", zap.Error(err))
					continue
				}
				if !ev.Interesting() {
					continue
				}
				if err := ce.Add(ev.ActorID, ev.Actor); err != nil {
					return errors.Wrap(err, "add")
				}
				{
					colEv.Append(proto.Enum8(ev.Type))
					colRepoID.Append(ev.RepoID)
					colActorID.Append(ev.ActorID)
					colTime.Append(ev.Time)
				}
			}
			if err := s.Err(); err != nil {
				return errors.Wrap(err, "scan")
			}
			// Flush remaining data.
			if err := write(); err != nil {
				return errors.Wrap(err, "write")
			}
			if err := outWrite.Flush(); err != nil {
				return errors.Wrap(err, "flush")
			}
			if err := outWrite.Close(); err != nil {
				return errors.Wrap(err, "close")
			}
			if err := outFile.Sync(); err != nil {
				return errors.Wrap(err, "sync")
			}
			if err := outFile.Close(); err != nil {
				return errors.Wrap(err, "close")
			}
			if err := os.Rename(outPathTmp, outPathTarget); err != nil {
				return errors.Wrap(err, "rename")
			}
			return nil
		})

		return g.Wait()
	})
}
