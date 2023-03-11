package archive

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	_testV = []byte("ernado")
	_testK = int64(866677)
)

func path(tb testing.TB) string {
	return filepath.Join(tb.TempDir(), "archive.bbolt")
}

func opts(tb testing.TB) UserCacheOptions {
	return UserCacheOptions{
		Path: path(tb),
		Size: 10_000,
	}
}

func newUserCache(tb testing.TB) *UserCache {
	tb.Helper()
	uc, err := NewUserCache(opts(tb))
	require.NoErrorf(tb, err, "NewUserCache() failed: %v", err)
	tb.Cleanup(func() {
		_ = uc.Close()
	})
	return uc
}

func TestUserCache(t *testing.T) {
	uc := newUserCache(t)

	require.NoErrorf(t, uc.Add(_testK, _testV), "Add()")
	require.True(t, uc.lru.Contains(_testK), "lru.Contains() failed")

	has, err := uc.inDB(uc.key(_testK))
	require.NoErrorf(t, err, "inDB() failed: %v", err)
	require.True(t, has, "inDB() failed")
}

func BenchmarkUserCache_Add(b *testing.B) {
	b.ReportAllocs()

	uc := newUserCache(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		require.NoErrorf(b, uc.Add(_testK, _testV), "Add()")
	}
}

func BenchmarkUserCacheWrite(b *testing.B) {
	b.ReportAllocs()
	uc := newUserCache(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		require.NoErrorf(b, uc.put(uc.key(_testK), _testV), "put()")
	}
}

func BenchmarkUserCacheRead(b *testing.B) {
	b.ReportAllocs()
	uc := newUserCache(b)
	require.NoErrorf(b, uc.put(uc.key(_testK), _testV), "put()")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		has, err := uc.inDB(uc.key(_testK))
		require.NoErrorf(b, err, "inDB() failed: %v", err)
		require.True(b, has, "inDB() failed")
	}
}
