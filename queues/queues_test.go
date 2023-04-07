package queue

import (
	"testing"
)

var capacity int
var queues [2]IQueue

func createQueues(capacity int) {
	arrayQueue := &ArrayQueue{}
	(*arrayQueue).Init(capacity)
	queues = [2]IQueue{arrayQueue, &LinkedQueue{}}
}

func deleteQueues() {
	queues[0] = nil
	queues[1] = nil
}

func setupTest() func() {
	capacity = 10
	createQueues(capacity)

	return func() {
		deleteQueues()
	}
}

func TestEnqueue(t *testing.T) {
	defer setupTest()()

	for _, queue := range queues {
		for i := 0; i < capacity; i++ {
			queue.Enqueue(i)
			if queue.Size() != i+1 {
				t.Errorf("%T size is %d, but we expected it to be %d", queue, queue.Size(), i+1)
			}
		}
	}
}

func TestDenqueue(t *testing.T) {
	defer setupTest()()

	for _, queue := range queues {
		for i := 0; i < capacity; i++ {
			queue.Enqueue(i)
		}

		element, _ := queue.Peek()
		if element != 0 {
			t.Errorf("%T returnned %d, but we expected it to be %d", queue, element, 0)
		}

		queue.Dequeue()
		element, _ = queue.Peek()
		if element != 1 {
			t.Errorf("%T returnned %d, but we expected it to be %d", queue, element, 1)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	defer setupTest()()

	for _, queue := range queues {
		if !queue.IsEmpty() {
			t.Errorf("we expected the %T to be empty but it returned false", queue)
		}
	}
}
