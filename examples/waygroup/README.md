# WaitGroup package example loop

This example shows how to use the WaitGroup package to implement a simple for loop.

The example is implemented in the file `waygroup.go`.


## Code overview

The example is implemented in the file `waygroup_test.go`.

```go
package main

import (
	"sync"
	"testing"
)

func BenchmarkMutexCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var counter int64
		var wg sync.WaitGroup
		var mu sync.Mutex
		ch := make(chan int64)

		// Launch 10 goroutines to increment the counter
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				for j := 0; j < 1000; j++ {
					mu.Lock()
					counter++
					mu.Unlock()
				}
				ch <- 1
				wg.Done()
			}()
		}

		// Wait for all goroutines to complete
		go func() {
			wg.Wait()
			close(ch)
		}()

		// Receive from the channel until it's closed
		for v := range ch {
			_ = v
		}
	}
}

```

## Running the example

To run the example, simply type:

```bash
go test -bench=. -benchmem

goos: linux
goarch: amd64
pkg: gobyexample/examples/waygroup
cpu: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz
BenchmarkMutexCounter-2             3782            318478 ns/op             636 B/op         15 allocs/op
PASS
ok      gobyexample/examples/waygroup   1.242s
```
