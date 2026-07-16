package main

import (
	"fmt"
	"slices"
)

func main() {

	//For slices you don't define the number of elements only the type

	var s []string
	fmt.Println("unitilised: ", s, s == nil, len(s) == 0)

	// use make to define a slice with a length
	s = make([]string, 3)
	fmt.Println("capacity: ", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s[2])

	//append items
	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println(s)

	//copy a slice
	c := make([]string, len(s))
	copy(c, s)

	fmt.Println(c)

	//slice operators
	fmt.Println(s[:3])
	fmt.Println(s[:1])
	fmt.Println(s[2:4])

	//init in one line
	d := []int{1, 2, 3}
	fmt.Println(d)

	e := []int{1, 2, 3}
	fmt.Println("Are the slices equal:", slices.Equal(d, e))

	//Other operaters
	//slices.Insert()

	toChunk := []int{1, 2, 3, 4, 5, 6, 7, 8}
	chunked := slices.Chunk(toChunk, 3)

	for i := range chunked {
		fmt.Println(i)
	}
}
