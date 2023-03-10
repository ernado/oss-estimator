package gh

import (
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/v50/github"
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

// Client initializes and returns new GitHub client from environmen
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
