package hashbenchmark

import (
	"fmt"
	"runtime"
	"testing"

	rand4 "github.com/cespare/xxhash/v2"
	rand3 "github.com/zeebo/wyhash"
	rand2 "github.com/zeebo/xxh3"
	rand1 "github.com/zhangyunhao116/wyhash"
)

func BenchmarkAll(b *testing.B) {
	sizes := []int{3, 6, 9, 12, 15, 21,
		24, 29, 32, 35, 64, 69, 96, 97,
		128, 129, 240, 241, 512, 1024, 100 * 1024,
	}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("rand1-%d", size), func(b *testing.B) {
			b.SetBytes(int64(size))
			var acc uint64
			d := string(make([]byte, size))
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				acc = rand1.Sum64String(d)
			}
			runtime.KeepAlive(acc)
		})
		b.Run(fmt.Sprintf("rand2-%d", size), func(b *testing.B) {
			b.SetBytes(int64(size))
			var acc uint64
			d := string(make([]byte, size))
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				acc = rand2.HashString(d)
			}
			runtime.KeepAlive(acc)
		})
		b.Run(fmt.Sprintf("rand3-%d", size), func(b *testing.B) {
			b.SetBytes(int64(size))
			var acc uint64
			d := string(make([]byte, size))
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				acc = rand3.HashString(d, 0)
			}
			runtime.KeepAlive(acc)
		})

		b.Run(fmt.Sprintf("rand4-%d", size), func(b *testing.B) {
			b.SetBytes(int64(size))
			var acc uint64
			d := string(make([]byte, size))
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				acc = rand4.Sum64String(d)
			}
			runtime.KeepAlive(acc)
		})
	}
}
