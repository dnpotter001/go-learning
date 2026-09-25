package main

import "fmt"

type base struct {
	num int
}

//we are able to embed one struct inside another
//different to inheritance because there is no polymorphism. There is no means of substitution, no "this is a ..." relationship

type container struct {
	base //embedded struct looks like a field without a type
	str  string
}

func (b base) describe() string { //this functions can be called on types of base
	return fmt.Sprintf("base with num= %v", b.num)
}

func main() {
	co := container{ //this is composition over inheritance
		base: base{
			num: 12,
		},
		str: "hello from the container",
	}

	fmt.Println(co)

	fmt.Println(co.base.num)

	fmt.Println("describe:", co.describe())      //we are able to call describe on the container.
	fmt.Println("describe:", co.base.describe()) //... but also on the base
}
