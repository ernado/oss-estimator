package af

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseAffiliationsFile(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("_testdata", "sample.txt"))
	require.NoError(t, err)

	var count int
	fn := func(u User) error {
		t.Logf("%+v", u)
		count++
		return nil
	}
	require.NoError(t, Parse(bytes.NewReader(data), fn))
	require.Equal(t, 3, count)
}
