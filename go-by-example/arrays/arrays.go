package main

import (
	"fmt"
)

func main() {

	//slices are more common than arrays

	//define one
	var a [5]int
	a[0] = 10

	fmt.Println(a[1])   // nothing is assigned so we get the default value
	fmt.Println(len(a)) //built in length function

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b) //[1 2 3 4 5]

	// compiler counts the number of items
	c := [...]int{1, 2, 3, 4}
	fmt.Println(c)

	// define on item, the others are defaulted
	d := [...]int{1, 5: 2}
	fmt.Println(d)

	// two dimensional array
	var twoD [2][3]int

	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("two Dimenions", twoD)

	// creating and initialising at one
	twoD = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("two Dimenions", twoD[1]) // only the second row

}
