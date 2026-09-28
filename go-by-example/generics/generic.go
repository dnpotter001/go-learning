package main

import "fmt"

// takes a slice of e where e is comparable,
func SliceIndex[S ~[]E, E comparable](slice S, value E) int {
	for i := range slice {
		if value == slice[i] {
			return i
		}
	}

	return -1
}

// generic implementation of a singularly linked list
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (list *List[T]) Push(v T) {
	if list.tail == nil { //case for an empty list
		list.head = &element[T]{val: v}
		list.tail = list.head
	} else {
		list.tail.next = &element[T]{val: v}
		list.tail = list.tail.next
	}
}

func (list *List[T]) AllElements() []T {
	var elems []T
	for e := list.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	fmt.Println("index of zoo: ", SliceIndex(s, "zoo"))
	// ^ type is inferred
	// or can be defined below
	_ = SliceIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
