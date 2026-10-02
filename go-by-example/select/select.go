package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for range 2 { //runs twice, first run is channel one and 2nd is chanel two
		select { //we use the select to await the values on both channels
		// a select need at least one case to do something meaningful
		case msg1 := <-c1:
			fmt.Println("received", msg1)

		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
		//default: can use a default case for polling, it will run until a channel is updated
	}

	fmt.Println("Both channel awaited")
}
