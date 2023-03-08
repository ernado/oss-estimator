package entry

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

type EventType byte

// "WatchEvent", "PushEvent", "IssuesEvent", "PullRequestEvent"
const (
	EventUnknown EventType = iota
	EventWatch
	EventPush
	EventIssue
	EventPR
)

func (t EventType) String() string {
	switch t {
	case EventWatch:
		return "WatchEvent"
	case EventPush:
		return "PushEvent"
	case EventIssue:
		return "IssuesEvent"
	case EventPR:
		return "PullRequestEvent"
	case EventUnknown:
		return "Unknown"
	default:
		return fmt.Sprintf("Event(%d)", t)
	}
}

type Event struct {
	Type    EventType
	RepoID  int64
	ActorID int64
	Actor   []byte
	Time    time.Time
}

func (e Event) String() string {
	return fmt.Sprintf("%s %5s %d",
		e.Time.Format(Layout), e.Type, e.RepoID,
	)
}

func (e *Event) Reset() {
	e.Type = EventUnknown
	e.Time = time.Time{}
	e.Actor = e.Actor[:0]
	e.ActorID = 0
	e.RepoID = 0
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
					case "id":
						id, err := d.Num()
						if err != nil {
							return errors.Wrap(err, "id parse")
						}
						v, err := id.Int64()
						if err != nil {
							return errors.Wrap(err, "id")
						}
						e.ActorID = v
						return nil
					case "login":
						v, err := d.StrAppend(e.Actor[:0])
						if err != nil {
							return errors.Wrap(err, "actor")
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
			v, err := d.StrBytes()
			if err != nil {
				return errors.Wrap(err, "type")
			}
			switch string(v) {
			case "WatchEvent":
				e.Type = EventWatch
			case "PushEvent":
				e.Type = EventPush
			case "IssuesEvent":
				e.Type = EventIssue
			case "PullRequestEvent":
				e.Type = EventPR
			default:
				e.Type = EventUnknown
			}
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
				case "id":
					id, err := d.Int64()
					if err != nil {
						return errors.Wrap(err, "id parse")
					}
					e.RepoID = id
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
	return nil
}

func (e Event) Full() bool {
	if len(e.Actor) == 0 {
		return false
	}
	if e.Time.IsZero() {
		return false
	}
	if e.Type == EventUnknown {
		return false
	}
	if e.ActorID == 0 {
		return false
	}
	if e.RepoID == 0 {
		return false
	}
	return true
}

func (e Event) Interesting() bool {
	return e.Type != EventUnknown
}
