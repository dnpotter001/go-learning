package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums { //range over an array or slice gives us an index and value
		sum += num
	}

	fmt.Println(sum)

	dictionary := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range dictionary {
		fmt.Println("%s -> %s", k, v)
	}

}
