package archive

import (
	"testing"
	"time"
)

func TestGetURL(t *testing.T) {
	t.Log(GetURL(time.Now().AddDate(0, 0, -2)))
}
