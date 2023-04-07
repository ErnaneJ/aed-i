package queue

type IQueue interface {
	Enqueue(int)
	Dequeue() (int, error)
	Peek() (int, error)
	Size() int
	IsEmpty() bool
}
