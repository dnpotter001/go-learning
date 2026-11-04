package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) { //one send channel and one receive channel
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished", j)
		results <- j * 2
	}
}

func main() {

	const jobNumbers int = 5
	jobs := make(chan int, jobNumbers)
	results := make(chan int, jobNumbers)

	var wg sync.WaitGroup

	for i := range 3 {
		wg.Go(func() { worker(i, jobs, results) })
	}

	for j := range 5 {
		jobs <- j * 2
	}
	close(jobs)
	// without this, range jobs in each worker never knows to stop, so workers never return and wg.Wait() blocks forever
	// an empty buffer blocks and waits, range jobs doesn't exit it just pauses, assuming there are more values
	// once the chanel is close range jobs drains teh remaining buffer values. then exits the loops
	wg.Wait()

}
