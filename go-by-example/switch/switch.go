package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now().Hour()

	//this switch doesn't switch on a value and has a default case
	switch {
	case t < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's afternoon")
	}

	//you can switch on type
	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case string:
			fmt.Println("I'm a string")
		default:
			fmt.Println("I'm something else")
		}
	}

	whatAmI(true)
	whatAmI("Hello world")
	whatAmI(123)

}
