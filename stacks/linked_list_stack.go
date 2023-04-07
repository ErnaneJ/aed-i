package stack

import (
	"errors"
)

type LinkedListStack struct {
	head   *Node
	length int
}

type Node struct {
	value int
	next  *Node
}

func (linkedListStack *LinkedListStack) Push(value int) {
	linkedListStack.head = &Node{value: value, next: linkedListStack.head}
	linkedListStack.length++
}

func (linkedListStack *LinkedListStack) Pop() (int, error) {
	if linkedListStack.length > 0 {
		item := linkedListStack.head.value
		linkedListStack.head = linkedListStack.head.next
		linkedListStack.length--

		return item, nil
	} else {
		return -1, errors.New("linkedListStack is empity")
	}
}

func (linkedListStack *LinkedListStack) Peek() (int, error) {
	if linkedListStack.length > 0 {
		return linkedListStack.head.value, nil
	} else {
		return -1, errors.New("linkedListStack is empity")
	}
}

func (linkedListStack *LinkedListStack) Empty() bool {
	return linkedListStack.length == 0
}

func (linkedListStack *LinkedListStack) Size() int {
	return linkedListStack.length
}
