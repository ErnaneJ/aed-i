package tree

type BinarySearchTree struct {
	value       int
	left, right *BinarySearchTree
}

type ITree interface {
	Insert(parent *BinarySearchTree, value int)
	Remove(value int) *BinarySearchTree
	Search(value int) bool
	Height() int
	Max() int
	Min() int
	IsBST() bool
	Size() int
	Print(space int)
}
