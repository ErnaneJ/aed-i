package tree

import "fmt"

func (AVLRoot *AVL) RotateLeft() *AVL {
	right := AVLRoot.right
	AVLRoot.right = right.left
	right.left = AVLRoot

	AVLRoot.UpdateProperties()
	right.UpdateProperties()

	return right
}

func (AVLRoot *AVL) RotateRight() *AVL {
	left := AVLRoot.left
	AVLRoot.left = left.right
	left.right = AVLRoot

	AVLRoot.UpdateProperties()
	left.UpdateProperties()

	return left
}

func (AVLRoot *AVL) Remove(value int) *AVL {
	if value < AVLRoot.value {
		AVLRoot.left = AVLRoot.left.Remove(value)
	} else if value > AVLRoot.value {
		AVLRoot.right = AVLRoot.right.Remove(value)
	} else { // BinarySearchTree found
		if AVLRoot.left == nil && AVLRoot.right == nil { // primary case: node without children
			return nil
		} else if AVLRoot.left != nil && AVLRoot.right == nil { // second case - left: node with 1 child
			return AVLRoot.left
		} else if AVLRoot.left == nil && AVLRoot.right != nil { // second case - right: node with 1 child
			return AVLRoot.right
		} else { // third case: node with 2 children
			maxNodeInLeft := AVLRoot.left.Max()
			AVLRoot.value = maxNodeInLeft
			AVLRoot.left = AVLRoot.left.Remove(maxNodeInLeft)
			return AVLRoot
		}
	}

	return AVLRoot
}

func (AVLRoot *AVL) Max() int {
	if AVLRoot == nil {
		return -1
	} else {
		if AVLRoot.right == nil {
			return AVLRoot.value
		} else {
			return AVLRoot.right.Max()
		}
	}
}

func (AVLRoot *AVL) Insert(value int) {
	if value <= AVLRoot.value {
		if AVLRoot.left == nil {
			AVLRoot.left = &AVL{value: value}
		} else {
			AVLRoot.left.Insert(value)
		}
	} else {
		if AVLRoot.right == nil {
			AVLRoot.right = &AVL{value: value}
		} else {
			AVLRoot.right.Insert(value)
		}
	}

	AVLRoot.UpdateProperties()
}

func (AVLRoot *AVL) Print(space int) {
	if AVLRoot == nil {
		return
	}

	const increment = 5

	space += increment
	AVLRoot.right.Print(space)
	fmt.Printf("%*s%d\n", space, "", AVLRoot.value)
	AVLRoot.left.Print(space)
}

func (AVLRoot *AVL) UpdateProperties() {
	heightRight := 0
	heightLeft := 0
	if AVLRoot.right == nil && AVLRoot.left == nil {
		AVLRoot.height = 0
	} else {
		if AVLRoot.right != nil {
			heightRight = AVLRoot.right.height
		}
		if AVLRoot.left != nil {
			heightLeft = AVLRoot.left.height
		}
		if heightRight > heightLeft {
			AVLRoot.height = 1 + heightRight
		} else {
			AVLRoot.height = 1 + heightLeft
		}
	}

	AVLRoot.bf = heightRight - heightLeft
}

func (AVLRoot *AVL) Search(value int) bool {
	if AVLRoot == nil {
		return false
	} else {
		if value < AVLRoot.value {
			return AVLRoot.left.Search(value)
		} else if value > AVLRoot.value {
			return AVLRoot.right.Search(value)
		} else {
			return true // value == AVLRoot.value
		}
	}
}

func (AVLRoot *AVL) Min() int {
	if AVLRoot == nil {
		return -1
	} else {
		if AVLRoot.left == nil {
			return AVLRoot.value
		} else {
			return AVLRoot.left.Min()
		}
	}
}

func (AVLRoot *AVL) Size() int {
	if AVLRoot == nil {
		return 0
	} else {
		return 1 + AVLRoot.left.Size() + AVLRoot.right.Size()
	}
}
