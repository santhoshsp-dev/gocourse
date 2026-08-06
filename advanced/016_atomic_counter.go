package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Step: 1)
type AtomicCounter struct {
	count int64
}

func (ac *AtomicCounter) increment() {
	atomic.AddInt64(&ac.count, 1)
}

func (ac *AtomicCounter) getValue() int64 {
	return atomic.LoadInt64(&ac.count)
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 10
	counter := &AtomicCounter{} // Step: 1)
	// value := 0 // Step: 2) if we are using Step: 2 we need to comment Step: 1 code

	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment() // Step: 1)
				// value++ // Step: 2) if we are using Step: 2 we need to comment Step: 1 code
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter.getValue()) // Step: 1)
	// fmt.Printf("Final counter value: %d\n", value) // Step: 2) if we are using Step: 2 we need to comment Step: 1 code
}
