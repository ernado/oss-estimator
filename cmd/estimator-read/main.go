package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"time"

	"github.com/ClickHouse/ch-go"
	"github.com/ClickHouse/ch-go/proto"
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.uber.org/zap"

	"estimator/internal/app"
)

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		var arg struct {
			Data string
			DB   string
		}
		flag.StringVar(&arg.DB, "db", "", "path to bbolt user database")
		flag.StringVar(&arg.Data, "data", "_data", "data directory")
		flag.Parse()

		// Load affiliations.
		var cncfRepos proto.ColInt64
		{
			data, err := os.ReadFile(filepath.Join(arg.Data, "cncf-repos.json"))
			if err != nil {
				return errors.Wrap(err, "read cncf repos")
			}
			if err := json.Unmarshal(data, &cncfRepos); err != nil {
				return errors.Wrap(err, "unmarshal cncf repos")
			}
		}

		af := make(map[int64]struct{})
		f, err := os.Open(filepath.Join(arg.Data, "affiliations.jsonl"))
		if err != nil {
			return errors.Wrap(err, "open affiliations")
		}
		defer func() { _ = f.Close() }()

		s := bufio.NewScanner(f)
		var d jx.Decoder
		for s.Scan() {
			d.ResetBytes(s.Bytes())
			var id int64
			if err := d.ObjBytes(func(d *jx.Decoder, key []byte) error {
				switch string(key) {
				case "id":
					v, err := d.Int64()
					if err != nil {
						return errors.Wrap(err, "id")
					}
					id = v
					return nil
				default:
					return d.Skip()
				}
			}); err != nil {
				return errors.Wrap(err, "decode")
			}
			af[id] = struct{}{}
		}
		if err := s.Err(); err != nil {
			return errors.Wrap(err, "scan")
		}

		c, err := ch.Dial(ctx, ch.Options{
			Database:    "estimator",
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

			total  int
			hit    int
			blocks int
		)
		start := time.Now()

		if err := c.Do(ctx, ch.Query{
			Body: "SELECT * FROM events WHERE repo IN (SELECT id FROM _data) AND event in ('IssuesEvent', 'PushEvent')",
			ExternalData: []proto.InputColumn{
				{Name: "id", Data: cncfRepos},
			},
			OnResult: func(ctx context.Context, block proto.Block) error {
				for i := 0; i < colActorID.Rows(); i++ {
					if _, ok := af[colActorID.Row(i)]; ok {
						hit++
					}
				}
				total += block.Rows
				blocks += 1
				if blocks%500 == 0 {
					lg.Info("Progress",
						zap.Int64("rps", int64(float64(total)/time.Since(start).Seconds())),
						zap.Int("total", total),
						zap.Int("hit", hit),
						zap.Int("miss", total-hit),
					)
				}
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

		lg.Info("Done",
			zap.Int64("rps", int64(float64(total)/time.Since(start).Seconds())),
			zap.Int("total", total),
			zap.Int("hit", hit),
			zap.Int("miss", total-hit),
		)

		return nil
	})
}
