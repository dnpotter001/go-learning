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

	// this is an unbuffered channel, it can't hold a value, the item has to be received after it is sent to avoid deadlock
	jobs := make(chan int)
	results := make(chan int)

	for i := range 3 { //fire off three workers
		go worker(i, jobs, results) //this starts three workers but they are all blocked because the jobs are empty
	}

	go func() { //start a new go routine that send the job values to the workers
		for j := range 5 {
			jobs <- j * 2 // blocks here until a worker is ready to receive this value — no buffer means no queuing
		}
		close(jobs) //this indicates that there is nothing else coming
	}()

	//without this gorountine the main process would send all of the 5 jobs before receiving any of the 5 results. this would block the main goroutine and create a deadlock

	for range 5 {
		<-results // blocks here until some worker finishes and sends its result back
	}

}
