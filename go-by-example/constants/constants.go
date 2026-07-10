package main

import (
	"fmt"
	"math"
)

const name string = "dave"

func main() {

	//Go supports constants of character, string, boolean, and numeric values.
	fmt.Println(name)

	const n = 50000000
	const d = 3e20

	//Below gives error: ./constants.go:16:2: cannot assign to d (neither addressable nor a map index expression)
	//d = 2e20

	fmt.Println(n / d)

	fmt.Println(math.Sin(n))

}
