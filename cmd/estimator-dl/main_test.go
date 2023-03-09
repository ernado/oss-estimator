package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	_testV = []byte("ernado")
	_testK = int64(866677)
)

func TestUserCache(t *testing.T) {
	uc, err := NewUserCache(t.TempDir(), 10_000)
	require.NoErrorf(t, err, "NewUserCache() failed: %v", err)
	defer func() {
		require.NoErrorf(t, uc.Close(), "Close() failed: %v", err)
	}()

	require.NoErrorf(t, uc.Add(_testK, _testV), "Add() failed: %v", err)
	require.True(t, uc.lru.Contains(_testK), "lru.Contains() failed")

	has, err := uc.inDB(uc.key(_testK))
	require.NoErrorf(t, err, "inDB() failed: %v", err)
	require.True(t, has, "inDB() failed")
}

func BenchmarkUserCache_Add(b *testing.B) {
	b.ReportAllocs()

	uc, err := NewUserCache(b.TempDir(), 10_000)
	require.NoErrorf(b, err, "NewUserCache() failed: %v", err)
	defer func() {
		require.NoErrorf(b, uc.Close(), "Close() failed: %v", err)
	}()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		require.NoErrorf(b, uc.Add(_testK, _testV), "Add() failed: %v", err)
	}
}

func BenchmarkUserCacheWrite(b *testing.B) {
	b.ReportAllocs()

	uc, err := NewUserCache(b.TempDir(), 10_000)
	require.NoErrorf(b, err, "NewUserCache() failed: %v", err)
	defer func() {
		require.NoErrorf(b, uc.Close(), "Close() failed: %v", err)
	}()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		require.NoErrorf(b, uc.put(uc.key(_testK), _testV), "put() failed: %v", err)
	}
}

func BenchmarkUserCacheRead(b *testing.B) {
	b.ReportAllocs()

	uc, err := NewUserCache(b.TempDir(), 10_000)
	require.NoErrorf(b, err, "NewUserCache() failed: %v", err)
	defer func() {
		require.NoErrorf(b, uc.Close(), "Close() failed: %v", err)
	}()
	require.NoErrorf(b, uc.put(uc.key(_testK), _testV), "put() failed: %v", err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		has, err := uc.inDB(uc.key(_testK))
		require.NoErrorf(b, err, "inDB() failed: %v", err)
		require.True(b, has, "inDB() failed")
	}
}
