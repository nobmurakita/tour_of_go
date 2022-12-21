package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func NewList[T any](val T) *List[T] {
	return &List[T]{val: val}
}

func (list *List[T]) append(val T) {
	p := list
	for p.next != nil {
		p = p.next
	}
	p.next = &List[T]{val: val}
}

func main() {
	li := NewList(1)
	li.append(2)
	li.append(3)
	li.append(4)
	li.append(5)
	for p := li; p != nil; p = p.next {
		fmt.Println(p.val)
	}

	ls := NewList("a")
	ls.append("b")
	ls.append("c")
	ls.append("d")
	ls.append("e")
	for p := ls; p != nil; p = p.next {
		fmt.Println(p.val)
	}
}
