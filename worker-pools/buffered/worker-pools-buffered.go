package main

import (
	"fmt"
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
	// this is an buffered channel as it holds 5 values.
	jobs := make(chan int, jobNumbers) //this allows the channels to receive up to 5 values,
	results := make(chan int, jobNumbers)
	// Without the buffered channel this would cause a dead lock.
	// This is because main would be blocked on sending the 4th value (there are three workers)
	// The results are not being read by the main thread because this is blocked by it not sending to the jobs channel.
	// And the results are not being read by the main channel because it's waiting to send the next job value
	// The workers have nowhere to put their results because main is not reading the results yet, it is still trying to send job value 3
	// This creates a deadlock between the main goroutine and workers

	for i := range 3 { //fire off three workers
		go worker(i, jobs, results) //this starts three workers but they are all blocked because the jobs are empty
	}

	for j := range 5 {
		jobs <- j * 2 // we then send 5 things to the worker.
	}
	close(jobs) //this indicates that there is nothing else coming

	for range 5 {
		<-results
	}

}
