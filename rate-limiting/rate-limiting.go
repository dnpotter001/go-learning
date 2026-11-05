package main

import (
	"fmt"
	"time"
)

func main() {

	//fake incoming requests
	requests := make(chan int, 5) //buffer of 5
	for i := range 5 {
		requests <- i * 2
	}
	close(requests)

	// the limiter channel will receive a new tick every 1 second
	limiter := time.Tick(1 * time.Second)

	for request := range requests {
		<-limiter // main routine is blocked until a new limiter tick is delivered
		fmt.Println("Request value: ", request, time.Now())
	}

	// we can process burst of requests
	processingTokens := make(chan time.Time, 3) //buffer of three with type time.

	// token bucket pattern, this buffer simulates a full bucket
	for range 3 {
		processingTokens <- time.Now() //filling up the channel will represent the first burst of messages
	}

	go func() { //new go routine to simulate filling of the bucket
		for t := range time.Tick(1 * time.Second) {
			processingTokens <- t // every 1 second we try and add a new value to the burst limiter
		}
	}()

	burstRequests := make(chan int, 5) //another buffer of 5 to simulate incoming message
	for i := range 5 {
		burstRequests <- i * 3
	}
	close(burstRequests)

	for req := range burstRequests {
		<-processingTokens //when the burst limit becomes unblocks the request is processed
		fmt.Println("Burst request: ", req, time.Now())
	}

}
