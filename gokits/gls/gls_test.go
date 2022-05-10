package gls

import (
	"testing"
	"time"
)

func BenchmarkGetPut(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Put(1 * time.Millisecond)
			getStorage()
		}
	})
}

func BenchmarkGets(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		Put(30 * time.Second)
		for pb.Next() {
			getStorage()
		}
	})
}
