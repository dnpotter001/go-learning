package main

import "fmt"

func main() {

	i := 20

	if i%10 == 0 {
		fmt.Println("Multiple of 10")
	} else if i%5 == 0 {
		fmt.Println("This won't fire")
	} else {
		fmt.Println("Something else")
	}
}
