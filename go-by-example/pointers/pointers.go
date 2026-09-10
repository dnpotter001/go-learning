package main

import "fmt"

func zeroVal(ival int) { //this gets a copy of the pointer
	ival = 0
}

func zeroPointer(iPointer *int) {
	*iPointer = 10
}

func main() {
	i := 1
	fmt.Println("initial value: ", i)

	zeroVal(i)
	fmt.Println("zeroval:", i) //value is still 1

	zeroPointer(&i)
	fmt.Println("zeroptr:", i) //changes value of the pointer instead of copying the value
	// this allows you to have side effects as go passes everything by value

	fmt.Println("pointer:", &i)

	newPointer := new(40)
	fmt.Println("address:", newPointer)

	fmt.Println("value:", *newPointer)

}
