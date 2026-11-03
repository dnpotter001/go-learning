package main

import (
	"fmt"
	"time"
)

func waitAndReturnOverChannel(timeoutInSeconds int, returnChanel chan string) {
	time.Sleep(time.Duration(timeoutInSeconds) * time.Second)
	returnChanel <- fmt.Sprintf("Waited for %d seconds", timeoutInSeconds)
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go waitAndReturnOverChannel(2, channel1)
	go waitAndReturnOverChannel(3, channel2)

	for range 2 {
		select {
		case msg1 := <-channel1:
			fmt.Println(msg1)
		case msg2 := <-channel2:
			fmt.Println(msg2)
		case <-time.After(4 * time.Second):
			fmt.Println("timeout after 4 second") //waits for 3 second before running the select
		}
	}

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select { //select is used to await the channel values from the ticker, the ticker returns the time it ticks at
			case <-done: //not actually needed for this demo to work
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond) //We wait while the ticker runs three times
	ticker.Stop()
	//done <- true
	fmt.Println("Ticker stopped")

}
