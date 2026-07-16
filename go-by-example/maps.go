package main

import (
	"fmt"
	"maps"
)

func main() {

	//maps are the associative data type in go a bit like a dictionary or hash
	m := make(map[string]int)

	//set the values
	m["k1"] = 1
	m["k2"] = 2

	// print them
	fmt.Println("map: ", m)

	fmt.Println("len:", len(m))

	delete(m, "k1")
	fmt.Println("after delete: ", m)

	clear(m)
	fmt.Println("after clear", m)

	m["k3"] = 3

	k, isValue := m["k3"] //is value shoes if there is a value present
	fmt.Println("value of key", k)
	fmt.Println("value of prs", isValue)

	// can init at the same time
	n := map[string]int{"key1": 1, "key2": 2}
	fmt.Println(n)

	n2 := map[string]int{"key1": 1, "key2": 2}
	fmt.Println("are they equal", maps.Equal(n, n2))
}
