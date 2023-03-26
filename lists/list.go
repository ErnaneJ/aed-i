package list

type DoublyLinkedList struct {
	head   *NodeDoublyLinkedList
	tail   *NodeDoublyLinkedList
	length int
}

type NodeDoublyLinkedList struct {
	value int
	next  *NodeDoublyLinkedList
	prev  *NodeDoublyLinkedList
}

type LinkedList struct {
	head   *NodeLinkedList
	length int
}

type NodeLinkedList struct {
	value int
	next  *NodeLinkedList
}

type IList interface {
	Add(value int)
	AddOnIndex(value int, index int) error
	Remove() error
	RemoveOnIndex(index int) error
	Get(index int) (int, error)
	Set(value int, index int) error
	Size() int
	ShowList()
}
