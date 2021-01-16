package hashbenchmark

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/cespare/xxhash/v2"
	zeebowyhash "github.com/zeebo/wyhash"
	"github.com/zeebo/xxh3"
	"github.com/zhangyunhao116/wyhash"
)

func BenchmarkAll(b *testing.B) {
	sizes := []int{0, 3, 4, 6, 8, 9, 11, 13, 15, 16,
		17, 21, 24, 29, 32,
		33, 64, 69, 96, 97, 128, 129, 240, 241,
		512, 1024, 100 * 1024,
	}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("wyhash-v1-%d", size), func(b *testing.B) {
			b.SetBytes(int64(size))
			var acc uint64
			d := make([]byte, size)
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				acc = wyhash.Sum64(d)
			}
			runtime.KeepAlive(acc)
		})
		b.Run(fmt.Sprintf("wyhash-v2-%d", size), func(b *testing.B) {
			b.SetBytes(int64(size))
			var acc uint64
			d := make([]byte, size)
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				acc = zeebowyhash.Hash(d, 0)
			}
			runtime.KeepAlive(acc)
		})
		b.Run(fmt.Sprintf("xxhash-v3-%d", size), func(b *testing.B) {
			b.SetBytes(int64(size))
			var acc uint64
			d := make([]byte, size)
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				acc = xxh3.Hash(d)
			}
			runtime.KeepAlive(acc)
		})
		b.Run(fmt.Sprintf("xxha64-v1-%d", size), func(b *testing.B) {
			b.SetBytes(int64(size))
			var acc uint64
			d := make([]byte, size)
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				acc = xxhash.Sum64(d)
			}
			runtime.KeepAlive(acc)
		})
	}
}
