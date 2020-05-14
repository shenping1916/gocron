package gocron

import (
	"testing"
)

func BenchmarkClosedNumber(b *testing.B) {
	s := &spec{}
	array := []uint{0, 15, 55}

	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = s.ClosedNumber(array, 16)
		}
	})
}
