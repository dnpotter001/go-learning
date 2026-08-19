package main

import "fmt"

func allInts(a, b, c int) (int, int, int) {
	return a, b, c
}

func varadicExample(nums ...int) {
	//nums is a slice
	for i := range nums {
		println(i) //simpler than the fmt functions
	}
}

func main() {
	fmt.Println(allInts(1, 2, 3))
	varadicExample(1, 2, 3, 4, 5)

	//a varadic function without input creates nil not a empty slice.
	varadicExample()

	//sliceA := []int{1,2,3, 4}
	//varadicExample(sliceA) <- does not work even though a slice is used underneath
	// A new slice is allocated
}
