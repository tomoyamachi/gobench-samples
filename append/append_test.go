package gobench_samples

import (
	"fmt"
	"testing"
)

func BenchmarkAppend_AllocateEveryTime(b *testing.B) {
	base := []string{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base = append(base, fmt.Sprintf("no%d", i))
	}
}

func BenchmarkAppend_AllocateOnceIndex(b *testing.B) {
	base := make([]string, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base[i] = fmt.Sprintf("no%d", i)
	}
}

func BenchmarkAppend_AllocateOnceAppend(b *testing.B) {
	base := make([]string, 0, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base = append(base, fmt.Sprintf("no%d", i))
	}
}

// BenchmarkAppend_AllocateEveryTime-12            10000000               225 ns/op
// BenchmarkAppend_AllocateOnceIndex-12            20000000               105 ns/op
// BenchmarkAppend_AllocateOnceAppend-12           20000000               104 ns/op
