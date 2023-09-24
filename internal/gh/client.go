package gh

import (
	"context"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"github.com/google/go-github/v50/github"
	"go.uber.org/multierr"
	"go.uber.org/ratelimit"
	"golang.org/x/oauth2"
)

type tokList struct {
	Tokens []*oauth2.Token
	Rate   ratelimit.Limiter
}

func (t *tokList) Token() (*oauth2.Token, error) {
	t.Rate.Take()
	n := len(t.Tokens)
	i := rand.Intn(n)
	return t.Tokens[i], nil
}

func check(client *github.Client) error {
	_, _, err := client.Repositories.Get(context.Background(), "ernado", "oss-estimator")
	if err != nil {
		return err
	}
	return err
}

func Check() error {
	// Check all tokens.
	var out error
	for _, v := range strings.Split(os.Getenv("GITHUB_TOKENS"), ",") {
		tok := &oauth2.Token{
			AccessToken: v,
		}
		httpClient := &http.Client{
			Transport: &oauth2.Transport{
				Base:   http.DefaultTransport,
				Source: oauth2.StaticTokenSource(tok),
			},
		}
		client := github.NewClient(httpClient)
		if err := check(client); err != nil {
			out = multierr.Append(out, errors.Wrapf(err, "token %s", tok.AccessToken[len(tok.AccessToken)-8:]))
		}
	}
	return out
}

// Client initializes and returns new GitHub client from environment.
func Client() *github.Client {
	var source tokList
	for _, v := range strings.Split(os.Getenv("GITHUB_TOKENS"), ",") {
		source.Tokens = append(source.Tokens, &oauth2.Token{
			AccessToken: v,
		})
	}
	const singleTokenPerHour = 5_000
	source.Rate = ratelimit.New(
		singleTokenPerHour*len(source.Tokens),
		ratelimit.Per(time.Hour),
	)
	httpClient := &http.Client{
		Transport: &oauth2.Transport{
			Base:   http.DefaultTransport,
			Source: &source,
		},
	}
	return github.NewClient(httpClient)
}
