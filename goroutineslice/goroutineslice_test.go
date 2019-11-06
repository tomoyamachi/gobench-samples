package goroutineslice

import (
	"sync"
	"testing"

	"golang.org/x/sync/errgroup"
)

var sampleCount = 10000000

func BenchmarkAppendSlice_MutexLock(b *testing.B) {
	b.ResetTimer()
	var g errgroup.Group
	intSlice := make([]int, 0, sampleCount)
	mu := sync.RWMutex{}
	for i := 0; i < sampleCount; i++ {
		i := i
		g.Go(func() error {
			mu.Lock()
			intSlice = append(intSlice, i)
			mu.Unlock()
			return nil
		})
	}
	g.Wait()
}

func BenchmarkAppendSlice_Channel(b *testing.B) {
	b.ResetTimer()
	var g errgroup.Group
	intChan := make(chan int, sampleCount)
	for i := 0; i < sampleCount; i++ {
		i := i
		g.Go(func() error {
			intChan <- i
			return nil
		})
	}
	g.Wait()
	intSlice := make([]int, 0, sampleCount)
	for i := 0; i < sampleCount; i++ {
		select {
		case target := <-intChan:
			intSlice = append(intSlice, target)
		}
	}
}

// BenchmarkAppend_AllocateEveryTime-12            10000000               225 ns/op
// BenchmarkAppend_AllocateOnceIndex-12            20000000               105 ns/op
// BenchmarkAppend_AllocateOnceAppend-12           20000000               104 ns/op
