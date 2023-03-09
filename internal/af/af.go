package af

import (
	"bufio"
	"io"
	"strings"
	"time"

	"github.com/go-faster/errors"
)

type Company struct {
	Name  string
	From  time.Time
	Until time.Time
}

type User struct {
	Name      string
	Emails    []string
	Companies []Company
}

// Parse parses the affiliations file from io.Reader.
//
// See https://github.com/cncf/gitdm, developers_affiliations*.txt.
func Parse(in io.Reader, fn func(u User) error) error {
	scanner := bufio.NewScanner(in)

	var u User
	for scanner.Scan() {
		s := scanner.Text()
		if strings.HasPrefix(s, "#") || s == "" {
			continue
		}
		if strings.HasPrefix(s, "\t") || strings.HasPrefix(s, "\u00a0") {
			// Company affiliation.
			var company Company
			elems := strings.Split(s, " ")
			var lastToken string
			for i, e := range elems {
				switch e {
				case "from", "until":
					if lastToken == "" {
						company.Name = strings.Join(elems[0:i], " ")
					}
					lastToken = e
				default:
					t, err := time.Parse("2006-01-02", e)
					if err != nil {
						continue
					}
					switch lastToken {
					case "from":
						company.From = t
					case "until":
						company.Until = t
					default:
						company.From = t
						lastToken = "from"
					}
				}
			}
			if lastToken == "" {
				company.Name = s
			}
			company.Name = strings.TrimSpace(company.Name)
			u.Companies = append(u.Companies, company)
		} else {
			if u.Name != "" {
				if err := fn(u); err != nil {
					return errors.Wrap(err, "fn")
				}
				u = User{}
			}
			idx := strings.Index(s, ":")
			if idx == -1 {
				return errors.Errorf("invalid line: %q", s)
			}
			u.Name = s[:idx]
			for _, e := range strings.Split(s[idx+1:], " ") {
				e = strings.Trim(e, ",")
				e = strings.ReplaceAll(e, "!", "@")
				e = strings.TrimSpace(e)
				if e == "" {
					continue
				}
				u.Emails = append(u.Emails, e)
			}
		}
	}
	return nil
}
