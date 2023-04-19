package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func BenchmarkAtomic(b *testing.B) {
	var counter int64

	// Run the loop for the specified number of iterations
	for n := 0; n < b.N; n++ {
		// Launch 10 goroutines to increment the counter
		for i := 0; i < 10; i++ {
			go func() {
				for j := 0; j < 1000; j++ {
					atomic.AddInt64(&counter, 1)
				}
			}()
		}

		// Wait for all goroutines to complete
		for i := 0; i < 10; i++ {
			<-time.After(time.Millisecond * 100)
		}
	}
}
