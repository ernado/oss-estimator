package entry

import (
	"time"
)

// Layout of chunk key.
//
// Was adjusted to be the prefix of RFC3339.
const Layout = "2006-01-02T15"

// Start time of entries.
func Start() time.Time {
	return time.Date(2014, 1, 1, 0, 0, 0, 0, time.UTC)
}

// Delta between chunks or chunk size.
const Delta = time.Hour
