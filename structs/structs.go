package main

import "fmt"

type person struct {
	firstName  string
	secondName string
	age        int
}

// define a method on struct type
func (p *person) formatName() string { //this method is using a pointer receiver  that will stop copying of the value when the method is called
	return p.firstName + " " + p.secondName
}

func makePerson(name string) *person {
	p := person{firstName: name}
	p.age = 42
	return &p //returns the pointer to the person, this will be cleaned up by gc
}

func main() {
	fmt.Println(person{firstName: "david", age: 29})

	fmt.Println(person{firstName: "david"}) // age is given the default value

	fmt.Println(makePerson("david").formatName())

	fmt.Println((&person{"david", "potter", 0}).formatName())

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog) //can define a stuct anonymously

}
