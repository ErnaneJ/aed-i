package tree

type BinarySearchTree struct {
	value       int
	left, right *BinarySearchTree
}

type ITree interface {
	Insert(value int)
	Remove(value int) *BinarySearchTree
	Search(value int) bool
	Height() int
	Max() int
	Min() int
	IsBST() bool
	Size() int
	Print(space int)
}

type AVL struct {
	bf, height, value int
	left, right       *AVL
}

type IAVL interface {
	RotateLeft() *AVL
	RotateRight() *AVL
	Insert(value int)
	Remove(value int) *AVL

	Max() int
	Print(space int)
	Search(value int) bool
	Min() int
	Size() int

	UpdateProperties()
}
