package entry

import (
	"bytes"
	"context"
	_ "embed"
	"io"
	"testing"
	"time"

	"github.com/klauspost/compress/zstd"

	"estimator/internal/speed"
)

//go:embed _testdata/2019-11-27T03.first100.json.zst
var dataCompressed []byte

func BenchmarkReader_Decode(b *testing.B) {
	ctx := context.Background()
	b.SetBytes(int64(len(dataCompressed)))
	b.ReportAllocs()
	b.ResetTimer()

	start := time.Now()

	b.RunParallel(func(pb *testing.PB) {
		var (
			r  = NewReader(speed.NewMetric())
			br = bytes.NewReader(dataCompressed)
		)
		defer r.Close()

		for pb.Next() {
			br.Reset(dataCompressed)
			if err := r.Decode(ctx, Decode{Reader: br}); err != nil {
				b.Fatal(err)
			}
		}
	})

	duration := time.Since(start)
	uncompressed := float64(b.N*len(dataMulti)) / duration.Seconds()
	uncompressed /= 1024 * 1024 // MB
	b.ReportMetric(uncompressed, "(json)MB/s")
}

func BenchmarkZSTD(b *testing.B) {
	b.SetBytes(int64(len(dataCompressed)))
	b.ReportAllocs()
	b.ResetTimer()

	start := time.Now()

	b.RunParallel(func(pb *testing.PB) {
		z, err := zstd.NewReader(nil, zstd.WithDecoderConcurrency(1))
		if err != nil {
			panic(err)
		}

		r := bytes.NewReader(nil)
		defer z.Close()

		for pb.Next() {
			r.Reset(dataCompressed)
			if err := z.Reset(r); err != nil {
				b.Fatal(err)
			}
			if _, err := io.Copy(io.Discard, z); err != nil {
				b.Fatal(err)
			}
		}
	})

	duration := time.Since(start)
	uncompressed := float64(b.N*len(dataMulti)) / duration.Seconds()
	uncompressed /= 1024 * 1024 // MB
	b.ReportMetric(uncompressed, "(json)MB/s")
}
