package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var truthey bool = true
	fmt.Println(truthey)

	// use the := to assign and decare shorthand
	f := "this is a cool shorthand way to do it"
	fmt.Println(f)

	//default values of primitives
	fmt.Println("are the default values what's expected?")

	var booleanDefault bool
	var intDefault int
	var stringDefault string
	var doubleDefault float32

	fmt.Println(booleanDefault)
	fmt.Println(intDefault)
	fmt.Println(stringDefault)
	fmt.Println(doubleDefault)

	fmt.Println("Yes they are")

}
