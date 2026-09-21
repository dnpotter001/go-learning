package main

import "fmt"

func main() {

	done := make(chan bool)
	go func(chan bool) {
		fmt.Println("this is from within a go routine")
		done <- true
	}(done) //trigger anonymously

	<-done // blocking chanel is used to wait

	// channel buffer

	channelBuffer := make(chan string, 2)

	channelBuffer <- "first"
	channelBuffer <- "second"
	//channelBuffer <- "third" //error from teh third value

	fmt.Println(<-channelBuffer)
	fmt.Println(<-channelBuffer)
	//fmt.Println(<-channelBuffer) //error from third value

	sharedChan := make(chan bool, 2)

	go func(chan bool) {
		oneToTen := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for val := range oneToTen {
			fmt.Println(oneToTen[val])
		}
		sharedChan <- true
	}(sharedChan)

	go func(chan bool) {
		tenToTwenty := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
		for val := range tenToTwenty {
			fmt.Println(tenToTwenty[val])
		}
		sharedChan <- true
	}(sharedChan) //trigger anonymously

	<-sharedChan
	<-sharedChan
	fmt.Println("completed")
	// this produces 11-20 fist as within the go scheduler there is a
	// runnext property and this one is occupied by the go routine which was queue last
	// There is a run next for each logical processor in the go scheduler
	// initially it is one to ten, then this get bumped into the FIFO queue

}
