package archive

import (
	"fmt"
	"time"
)

// GetURL returns github archive URL for the given date.
func GetURL(date time.Time) string {
	// You can't describe 24-hour without leading zero.
	// Correct format is yyyy-MM-dd-H.
	return fmt.Sprintf("https://data.gharchive.org/%s-%d.json.gz",
		date.Format("2006-01-02"), date.Hour(),
	)
}
