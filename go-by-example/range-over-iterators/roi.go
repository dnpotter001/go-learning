package main

import (
	"fmt"
	"iter"
	"slices"
)

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// the type Seq is a func(yield func(T) bool) so swapping the method signature to what is shown below compiles
// func (lst *List[T]) All() func(yield func(T) bool) {
func (lst *List[T]) All() iter.Seq[T] { //special iter method signature shows that it returns an iterator
	return func(yield func(T) bool) { //this iterator functions takes another function called yield (go convention),
		// the bool is part of the go spec and lets the range operator when to stop iterating
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) { //when yield is false we return (end of sequence, or break in a loop)
				return
			}
		}
	}
}

func main() {

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	for e := range lst.All() { //since All returns an iterator we can use the range key word
		fmt.Println(e)
	}

	all := slices.Collect(lst.All()) //collect function puts all values into a slice
	fmt.Println("all:", all)
}
