package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var operationsCountAtomic atomic.Uint64 //we can use this atomic type to represent our counter.
	var operationsCount int
	var waitGroup sync.WaitGroup // this will help us wait for all gorountines to finish there work

	for range 50 { //firing off 50 go routines, each adding 1 to the operations counter 50 times
		waitGroup.Go(func() {
			for range 1000 {
				operationsCountAtomic.Add(1)
			}
		})
	}

	for range 50 { //firing off 50 go routines, each adding 1 to the operations counter 50 times
		waitGroup.Go(func() {
			for range 1000 {
				operationsCount++
			}
		})
	}

	waitGroup.Wait() //wait for all operations to be done

	fmt.Println("atomic operations", operationsCountAtomic.Load())
	fmt.Println("operations", operationsCount)
	//we get exactly 50000 every time, when using a non atomic int
	//we get a different value each time for the non atomic as the goroutines read and write different values (lots of race conditions)
	//when running with -race we get a race condition error too
}
