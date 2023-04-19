package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var counter int64

	// Launch 10 goroutines to increment the counter
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	// Wait for all goroutines to complete
	time.Sleep(time.Second)

	// Print the final value of the counter
	fmt.Println("Counter:", counter)
}
