package main

import "fmt"

func main() {

	for {
		fmt.Println("this is inside a for that doesn't stop unless there is a break")
		break
	}

	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}

	for n := 0; n < 3; n++ {
		fmt.Println(n)
	}

	//You can use 'range' do say do this n number of times
	j := 10
	counter := 1
	for range j {
		if counter%2 == 0 {
			counter++
			continue
		}
		fmt.Println(counter)
		counter++
	}

}
