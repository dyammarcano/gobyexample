package main

import (
	"fmt"
	"sync"
)

func main() {
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

	// Print the final value of the counter
	fmt.Println("Counter:", counter)
}
