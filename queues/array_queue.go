package queue

import (
	"errors"
)

type ArrayQueue struct {
	front    int
	back     int
	values   []int
	length   int
	capacity int
}

func (arrayQueue *ArrayQueue) Init(capacity int) error {
	if arrayQueue.length > 0 {
		return errors.New("arrayQueue already initialized")
	}

	if capacity > 0 {
		arrayQueue.values = make([]int, capacity)
		arrayQueue.front = -1
		arrayQueue.back = -1
		arrayQueue.length = 0
		arrayQueue.capacity = capacity
		return nil
	} else {
		return errors.New("can't initialize ArrayQueue with size <= 0")
	}
}

// func (arrayQueue *ArrayQueue) double() {
// 	doubledValues := make([]int, len(arrayQueue.values)*2)
// 	for i := 0; i < len(arrayQueue.values); i++ {
// 		doubledValues[i] = arrayQueue.values[i]
// 	}
// 	arrayQueue.values = doubledValues
// }

func (arrayQueue *ArrayQueue) Enqueue(value int) {
	// if arrayQueue.length == len(arrayQueue.values) { arrayQueue.double() }
	// if arrayQueue.length == arrayQueue.capacity {
	// 	return errors.New("arrayQueue is full")
	// }

	if arrayQueue.length == 0 {
		arrayQueue.front = 0
		arrayQueue.back = 0
		arrayQueue.values[0] = value
		arrayQueue.length = 1
	} else {
		arrayQueue.back = (arrayQueue.back + 1) % arrayQueue.capacity
		arrayQueue.values[arrayQueue.back] = value
		arrayQueue.length++
	}
}

func (arrayQueue *ArrayQueue) Dequeue() (int, error) {
	if arrayQueue.length <= 0 {
		return -1, errors.New("arrayQueue is empty")
	}

	aux := arrayQueue.values[arrayQueue.front]
	if arrayQueue.length == 1 {
		arrayQueue.front = -1
		arrayQueue.back = -1
		arrayQueue.length = 0
		return aux, nil
	}

	arrayQueue.front = (arrayQueue.front + 1) % arrayQueue.capacity
	arrayQueue.length--
	return aux, nil
}

func (arrayQueue *ArrayQueue) Peek() (int, error) {
	if arrayQueue.length <= 0 {
		return -1, errors.New("arrayQueue is empty")
	}

	return arrayQueue.values[arrayQueue.front], nil
}

func (arrayQueue *ArrayQueue) Size() int {
	return arrayQueue.length
}

func (arrayQueue *ArrayQueue) IsEmpty() bool {
	return arrayQueue.length == 0
}
