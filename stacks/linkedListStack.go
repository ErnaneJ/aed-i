package stack

import (
	"errors"
)

type LinkedListStack struct {
	values []int
	length int
}

func (linkedListStack *LinkedListStack) Init(size int) error {
	if linkedListStack.length > 0 {
		return errors.New("array stack already initialized")
	}

	if size > 0 {
		linkedListStack.values = make([]int, size)
		linkedListStack.length = 0
	} else {
		return errors.New("can't initialize LinkedListStack with size <= 0")
	}

	return nil
}

func (linkedListStack *LinkedListStack) double() {
	doubledValues := make([]int, len(linkedListStack.values)*2)
	for i := 0; i < len(linkedListStack.values); i++ {
		doubledValues[i] = linkedListStack.values[i]
	}

	linkedListStack.values = doubledValues
}

func (linkedListStack LinkedListStack) Push(value int) {
	if linkedListStack.length == len(linkedListStack.values) {
		linkedListStack.double()
	}

	linkedListStack.values[linkedListStack.length] = value
	linkedListStack.length++
}

func (linkedListStack *LinkedListStack) Pop() (int, error) {
	if linkedListStack.length > 0 {
		item := linkedListStack.values[linkedListStack.length]
		linkedListStack.length--

		return item, nil
	} else {
		return -1, errors.New("linkedListStack is empity")
	}
}

func (linkedListStack *LinkedListStack) Peek() (int, error) {
	if linkedListStack.length > 0 {
		return linkedListStack.values[linkedListStack.length], nil
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
