package entry

import (
	"bytes"
	"fmt"
	"time"
	"unsafe"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

type Event struct {
	Type    []byte
	Repo    []byte
	RepoURL []byte
	URL     []byte
	Actor   []byte
	Time    time.Time
}

func (e Event) String() string {
	return fmt.Sprintf("%s %30s %s",
		e.Time.Format(Layout), e.Type, e.Repo,
	)
}

func (e *Event) Reset() {
	e.Type = e.Type[:0]
	e.Repo = e.Repo[:0]
	e.RepoURL = e.RepoURL[:0]
	e.URL = e.URL[:0]
	e.Time = time.Time{}
	e.Actor = e.Actor[:0]
}

func (e *Event) Decode(d *jx.Decoder) error {
	if err := d.ObjBytes(func(d *jx.Decoder, key []byte) error {
		switch string(key) {
		case "actor":
			switch d.Next() {
			case jx.String:
				v, err := d.StrAppend(e.Actor[:0])
				if err != nil {
					return errors.Wrap(err, "actor")
				}
				e.Actor = v
			case jx.Null:
				return d.Skip()
			case jx.Object:
				if err := d.ObjBytes(func(d *jx.Decoder, key []byte) error {
					switch string(key) {
					case "login":
						v, err := d.StrAppend(e.Actor[:0])
						if err != nil {
							return errors.Wrap(err, "name")
						}
						e.Actor = v
						return nil
					default:
						if err := d.Skip(); err != nil {
							return errors.Wrap(err, "skip")
						}
						return nil
					}
				}); err != nil {
					return errors.Wrap(err, "actor")
				}
			default:
				return errors.Errorf("unexpected actor type %s", d.Next())
			}
			return nil
		case "type":
			v, err := d.StrAppend(e.Type[:0])
			if err != nil {
				return errors.Wrap(err, "type")
			}
			e.Type = v
			return nil
		case "url":
			v, err := d.StrAppend(e.URL[:0])
			if err != nil {
				return errors.Wrap(err, "url")
			}
			e.URL = v
			return nil
		case "created_at":
			v, err := d.StrBytes()
			if err != nil {
				return errors.Wrap(err, "created at")
			}
			t, err := time.Parse(time.RFC3339, *(*string)(unsafe.Pointer(&v)))
			if err != nil {
				return errors.Wrap(err, "bytes")
			}
			e.Time = t
			return nil
		case "repo", "repository":
			if err := d.ObjBytes(func(d *jx.Decoder, key []byte) error {
				switch string(key) {
				case "url":
					v, err := d.StrAppend(e.RepoURL[:0])
					if err != nil {
						return errors.Wrap(err, "name")
					}
					e.RepoURL = v
					return nil
				case "full_name":
					v, err := d.StrAppend(e.Repo[:0])
					if err != nil {
						return errors.Wrap(err, "name")
					}
					e.Repo = v
					return nil
				case "name":
					if len(e.Repo) != 0 {
						return d.Skip()
					}
					v, err := d.StrAppend(e.Repo[:0])
					if err != nil {
						return errors.Wrap(err, "name")
					}
					e.Repo = v
					return nil
				default:
					if err := d.Skip(); err != nil {
						return errors.Wrap(err, "skip")
					}
					return nil
				}
			}); err != nil {
				return errors.Wrap(err, "repo")
			}
			return nil
		default:
			if err := d.Skip(); err != nil {
				return errors.Wrap(err, "skip")
			}
			return nil
		}
	}); err != nil {
		return errors.Wrap(err, "object")
	}
	if len(e.Repo) == 0 {
		v := bytes.TrimPrefix(e.RepoURL, []byte("https://api.github.com/repos/"))
		e.Repo = append(e.Repo, v...)
	}
	return nil
}

func (e Event) Full() bool {
	if len(e.Repo) == 0 {
		return false
	}
	if len(e.Actor) == 0 {
		return false
	}
	if e.Time.IsZero() {
		return false
	}
	if len(e.Type) == 0 {
		return false
	}
	return true
}

func (e Event) Interesting() bool {
	switch string(e.Type) {
	case "WatchEvent", "PushEvent", "IssuesEvent", "PullRequestEvent":
		return true
	default:
		return false
	}
}
