package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/ClickHouse/ch-go"
	"github.com/ClickHouse/ch-go/proto"
	"github.com/go-faster/errors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"estimator/internal/app"
	"estimator/internal/archive"
	"estimator/internal/estimate"
)

type Event struct {
	Type        byte
	OrgID       int64
	OrgName     string
	RepoID      int64
	RepoName    string
	ActorID     int64
	ActorName   string
	CompanyName string
	Time        time.Time
}

type Company struct {
	Name    string `json:"name"`
	FromStr string `json:"from"`
	ToStr   string `json:"to"`

	From time.Time `json:"-"`
	To   time.Time `json:"-"`
}

func (c *Company) Parse() error {
	var err error
	if c.FromStr != "" {
		c.From, err = time.Parse("2006-01-02", c.FromStr)
		if err != nil {
			return errors.Wrap(err, "parse from")
		}
	}
	if c.ToStr != "" {
		c.To, err = time.Parse("2006-01-02", c.ToStr)
		if err != nil {
			return errors.Wrap(err, "parse to")
		}
	}
	return nil
}

func (c *Company) Actual(t time.Time) bool {
	if !c.From.IsZero() && c.From.After(t) {
		return false
	}
	if !c.To.IsZero() && c.To.Before(t) {
		return false
	}
	return true
}

type Affiliation struct {
	Name      string    `json:"name"`
	ID        int64     `json:"id"`
	Companies []Company `json:"companies"`
}

type Repo struct {
	OrgID    int64
	OrgName  string
	RepoID   int64
	RepoName string
}

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
		var ag estimate.Aggregated
		{
			data, err := os.ReadFile(filepath.Join(arg.Data, "aggregated.json"))
			if err != nil {
				return errors.Wrap(err, "read cncf repos")
			}
			if err := json.Unmarshal(data, &ag); err != nil {
				return errors.Wrap(err, "unmarshal cncf repos")
			}
		}
		// Key is repo id.
		repos := make(map[int64]Repo, 1_000)
		for _, org := range ag.Organizations {
			for _, repo := range org.Repos {
				repos[repo.ID] = Repo{
					OrgID:    org.ID,
					OrgName:  org.Name,
					RepoID:   repo.ID,
					RepoName: repo.Name,
				}
			}
		}
		af := make(map[int64]Affiliation, 40_000)
		f, err := os.Open(filepath.Join(arg.Data, "affiliations.jsonl"))
		if err != nil {
			return errors.Wrap(err, "open affiliations")
		}
		defer func() { _ = f.Close() }()

		s := bufio.NewScanner(f)
		for s.Scan() {
			var v Affiliation
			if err := json.Unmarshal(s.Bytes(), &v); err != nil {
				return errors.Wrap(err, "unmarshal affiliation")
			}
			for i, c := range v.Companies {
				if err := c.Parse(); err != nil {
					return errors.Wrap(err, "parse company")
				}
				v.Companies[i] = c
			}
			af[v.ID] = v
		}
		if err := s.Err(); err != nil {
			return errors.Wrap(err, "scan")
		}
		if len(af) < 10_000 {
			return errors.Errorf("too few affiliations: %d", len(af))
		}
		lg.Info("Loaded",
			zap.Int("affiliations", len(af)),
			zap.Int("repos", len(repos)),
			zap.Int("orgs", len(ag.Organizations)),
		)

		events := make(chan Event, 1000)
		g, ctx := errgroup.WithContext(ctx)
		g.Go(func() error {
			defer close(events)

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
						total++
						u, ok := af[colActorID.Row(i)]
						if !ok {
							continue
						}
						repo, ok := repos[colRepoID.Row(i)]
						if !ok {
							lg.Debug("Unknown repo", zap.Int64("id", colRepoID.Row(i)))
							return errors.Errorf("unknown repo %d", colRepoID.Row(i))
						}
						hit += 1
						t := colTime.Row(i)
						company := "Independent"
						for _, c := range u.Companies {
							if !c.Actual(t) {
								continue
							}
							company = c.Name
						}
						ev := Event{
							Type:        byte(colEv.Row(i)),
							RepoID:      colRepoID.Row(i),
							RepoName:    repo.RepoName,
							OrgID:       repo.OrgID,
							OrgName:     repo.OrgName,
							ActorID:     colActorID.Row(i),
							ActorName:   u.Name,
							CompanyName: company,
							Time:        t,
						}
						select {
						case <-ctx.Done():
							return ctx.Err()
						case events <- ev:
							continue
						}
					}
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
		g.Go(func() error {
			c, err := ch.Dial(ctx, ch.Options{
				Database:    "estimator",
				Compression: ch.CompressionLZ4,
			})
			if err != nil {
				return errors.Wrap(err, "dial")
			}
			defer func() { _ = c.Close() }()

			var (
				colEv        proto.ColEnum8
				colOrgID     proto.ColInt64
				colOrgName   proto.ColStr
				colRepoID    proto.ColInt64
				colRepoName  proto.ColStr
				colActorID   proto.ColInt64
				colActorName proto.ColStr
				colCompany   proto.ColStr
				colTime      proto.ColDateTime
			)
			input := proto.Input{
				{Name: "event", Data: proto.Wrap(&colEv, archive.EventDDL)},
				{Name: "org_id", Data: &colOrgID},
				{Name: "org_name", Data: &colOrgName},
				{Name: "repo_id", Data: &colRepoID},
				{Name: "repo_name", Data: &colRepoName},
				{Name: "actor_id", Data: &colActorID},
				{Name: "actor_name", Data: &colActorName},
				{Name: "company", Data: &colCompany},
				{Name: "time", Data: &colTime},
			}
			if err := c.Do(ctx, ch.Query{
				Body:  input.Into("cncf"),
				Input: input,
				OnInput: func(ctx context.Context) error {
					input.Reset()
				Fetch:
					for {
						select {
						case <-ctx.Done():
							return ctx.Err()
						case e, ok := <-events:
							if !ok {
								// Done.
								return io.EOF
							}

							colEv.Append(proto.Enum8(e.Type))
							colOrgID.Append(e.OrgID)
							colOrgName.Append(e.OrgName)
							colRepoID.Append(e.RepoID)
							colRepoName.Append(e.RepoName)
							colActorID.Append(e.ActorID)
							colActorName.Append(e.ActorName)
							colCompany.Append(e.CompanyName)
							colTime.Append(e.Time)

							if colEv.Rows() > 40_000 {
								break Fetch
							}
						}
					}
					return nil
				},
			}); err != nil {
				return errors.Wrap(err, "create table")
			}

			return nil
		})

		if err := g.Wait(); err != nil {
			return errors.Wrap(err, "group")
		}

		return nil
	})
}
