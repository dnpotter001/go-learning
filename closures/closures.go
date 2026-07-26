package main

import "fmt"

func closure() func() int {
	closeOver := 1

	return func() int {
		closeOver++
		return closeOver
	}
}

func main() {

	myReturnedFunction := closure() //anonymous function is returned here

	fmt.Println(myReturnedFunction())
	fmt.Println(myReturnedFunction())

	anotherVar := 0

	iterate := func() int {
		anotherVar++
		return anotherVar
	}

	fmt.Println(iterate())
	fmt.Println(iterate())
}
