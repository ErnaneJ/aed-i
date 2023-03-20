package stack

import (
	"errors"
)

type ArrayStack struct {
	values []int
	length int
}

func (arrayStack *ArrayStack) Init(size int) error {
	if arrayStack.length > 0 {
		return errors.New("array stack already initialized")
	}

	if size > 0 {
		arrayStack.values = make([]int, size)
		arrayStack.length = 0
	} else {
		return errors.New("can't initialize ArrayStack with size <= 0")
	}

	return nil
}

func (arrayStack *ArrayStack) double() {
	doubledValues := make([]int, len(arrayStack.values)*2)
	for i := 0; i < len(arrayStack.values); i++ {
		doubledValues[i] = arrayStack.values[i]
	}

	arrayStack.values = doubledValues
}

func (arrayStack ArrayStack) Push(value int) {
	if arrayStack.length == len(arrayStack.values) {
		arrayStack.double()
	}

	arrayStack.values[arrayStack.length] = value
	arrayStack.length++
}

func (arrayStack *ArrayStack) Pop() (int, error) {
	if arrayStack.length > 0 {
		item := arrayStack.values[arrayStack.length]
		arrayStack.length--

		return item, nil
	} else {
		return -1, errors.New("arrayStack is empity")
	}
}

func (arrayStack *ArrayStack) Peek() (int, error) {
	if arrayStack.length > 0 {
		return arrayStack.values[arrayStack.length], nil
	} else {
		return -1, errors.New("arrayStack is empity")
	}
}

func (arrayStack *ArrayStack) Empty() bool {
	return arrayStack.length == 0
}

func (arrayStack *ArrayStack) Size() int {
	return arrayStack.length
}
