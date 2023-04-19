# Atomic package example loop

This example shows how to use the atomic ackage to implement a simple for loop.

The example is implemented in the file `atomic_for.go`.


## Code overview

The example is implemented in the file `atomic_for_test.go`.

```go
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
```

## Running the example

To run the example, simply type:

```bash
go test -bench=. -benchmem

goos: windows
goarch: amd64
pkg: gobyexample/examples/atomic_for
cpu: Intel(R) Core(TM)2 Duo CPU     E8400  @ 3.00GHz
BenchmarkAtomic
BenchmarkAtomic-2              1        1173798100 ns/op
PASS
```
