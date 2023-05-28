package tree

import (
	"testing"
)

func TestInsert(t *testing.T) {
	bst := &BinarySearchTree{5, nil, nil}

	values := []int{3, 7, 2, 4, 6, 8}

	for index, value := range values {
		bst.Insert(&bst, value)
		length := bst.Length()
		if length != index+2 {
			t.Errorf("%T length is %d, but we expected it to be %d", bst, length, index+2)
		}
	}
}

func TestHeight(t *testing.T) {
	bst := &BinarySearchTree{5, nil, nil}

	values := []int{3, 7, 2, 4, 6, 8}

	for _, value := range values {
		bst.Insert(&bst, value)
	}

	if bst.Height() != 2 {
		t.Errorf("%T height is %d, but we expected it to be %d", bst, bst.Height(), 2)
	}
}

func TestSearch(t *testing.T) {
	bst := &BinarySearchTree{5, nil, nil}

	values := []int{3, 7, 2, 4, 6, 8}

	for _, value := range values {
		bst.Insert(&bst, value)

		if !bst.Search(value) {
			t.Errorf("%T does not have the value %d in its elements.", bst, value)
		}
	}

	if bst.Search(48521) {
		t.Errorf("%T does not have the value %d in its elements but said it does", bst, 48521)
	}
}

func TestRemove(t *testing.T) {
	bst := &BinarySearchTree{5, nil, nil}

	values := []int{3, 7, 2, 4, 6, 8}

	for _, value := range values {
		bst.Insert(&bst, value)
	}

	for index := len(values) - 2; index >= 0; index-- {
		bst.Remove(values[index])
		length := bst.Length()
		if length != index+2 {
			t.Errorf("%T length is %d, but we expected it to be %d", bst, length, index+2)
		}
	}
}

func TestMax(t *testing.T) {
	bst := &BinarySearchTree{5, nil, nil}

	values := []int{3, 7, 2, 4, 6, 8}

	for _, value := range values {
		bst.Insert(&bst, value)
	}

	if bst.Max() != 8 {
		t.Errorf("%T has the greatest element %d but returned %d", bst, 8, bst.Max())
	}
}

func TestMin(t *testing.T) {
	bst := &BinarySearchTree{5, nil, nil}

	values := []int{3, 7, 2, 4, 6, 8}

	for _, value := range values {
		bst.Insert(&bst, value)
	}

	if bst.Min() != 2 {
		t.Errorf("%T has the smallest element %d but returned %d", bst, 2, bst.Min())
	}
}

func TestIsBST(t *testing.T) {
	bst := &BinarySearchTree{5, nil, nil}

	values := []int{3, 7, 2, 4, 6, 8}

	for _, value := range values {
		bst.Insert(&bst, value)
	}

	if !bst.IsBST() {
		t.Errorf("%T is a BST but returned no", bst)
	}

	bst.left.left.value = 7452

	if bst.IsBST() {
		t.Errorf("%T is not a BST but returned that it is", bst)
	}
}
