package SLL

import (
	"fmt"
	methodsinterfaces "go_example_2/methods_interfaces"
)

type SingleLinkedList struct {
	head *methodsinterfaces.SLLNode
	tail *methodsinterfaces.SLLNode
}

func NewSingleLinkedList() *SingleLinkedList {
	return new(SingleLinkedList)
}

func (list *SingleLinkedList) Add(v int) {
	newNode := &methodsinterfaces.SLLNode{Value: v}
	if list.head == nil {
		list.head = newNode
	} else if list.head == list.tail {
		list.tail.Next = newNode
	} else if list.tail != nil {
		list.tail.Next = newNode
	}
	list.tail = newNode
}

func (list *SingleLinkedList) String() string {
	s := ""
	for n := list.head; n != nil; n = n.Next {
		s += fmt.Sprintf(" {%d} ", n.GetValue())
	}
	return s
}
