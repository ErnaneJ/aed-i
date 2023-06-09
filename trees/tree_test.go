package tree

import (
	"testing"
)

// ############################# BST TESTS #############################

func TestInsertBST(t *testing.T) {
	bst := &BinarySearchTree{5, nil, nil}

	values := []int{3, 7, 2, 4, 6, 8}

	for index, value := range values {
		bst.Insert(value)
		size := bst.Size()
		if size != index+2 {
			t.Errorf("%T size is %d, but we expected it to be %d", bst, size, index+2)
		}
	}
}

func TestHeightBST(t *testing.T) {
	bst := &BinarySearchTree{5, nil, nil}

	values := []int{3, 7, 2, 4, 6, 8}

	for _, value := range values {
		bst.Insert(value)
	}

	if bst.Height() != 2 {
		t.Errorf("%T height is %d, but we expected it to be %d", bst, bst.Height(), 2)
	}
}

func TestSearchBST(t *testing.T) {
	bst := &BinarySearchTree{5, nil, nil}

	values := []int{3, 7, 2, 4, 6, 8}

	for _, value := range values {
		bst.Insert(value)

		if !bst.Search(value) {
			t.Errorf("%T does not have the value %d in its elements.", bst, value)
		}
	}

	if bst.Search(48521) {
		t.Errorf("%T does not have the value %d in its elements but said it does", bst, 48521)
	}
}

func TestRemoveBST(t *testing.T) {
	bst := &BinarySearchTree{5, nil, nil}

	values := []int{3, 7, 2, 4, 6, 8}

	for _, value := range values {
		bst.Insert(value)
	}

	for index := len(values) - 2; index >= 0; index-- {
		bst.Remove(values[index])
		size := bst.Size()
		if size != index+2 {
			t.Errorf("%T size is %d, but we expected it to be %d", bst, size, index+2)
		}
	}
}

func TestMaxBST(t *testing.T) {
	bst := &BinarySearchTree{5, nil, nil}

	values := []int{3, 7, 2, 4, 6, 8}

	for _, value := range values {
		bst.Insert(value)
	}

	if bst.Max() != 8 {
		t.Errorf("%T has the greatest element %d but returned %d", bst, 8, bst.Max())
	}
}

func TestMinBST(t *testing.T) {
	bst := &BinarySearchTree{5, nil, nil}

	values := []int{3, 7, 2, 4, 6, 8}

	for _, value := range values {
		bst.Insert(value)
	}

	if bst.Min() != 2 {
		t.Errorf("%T has the smallest element %d but returned %d", bst, 2, bst.Min())
	}
}

func TestIsBST(t *testing.T) {
	bst := &BinarySearchTree{5, nil, nil}

	values := []int{3, 7, 2, 4, 6, 8}

	for _, value := range values {
		bst.Insert(value)
	}

	if !bst.IsBST() {
		t.Errorf("%T is a BST but returned no", bst)
	}

	bst.left.left.value = 7452

	if bst.IsBST() {
		t.Errorf("%T is not a BST but returned that it is", bst)
	}
}

func TestSortedArrayToBST(t *testing.T) {
	array := []int{50, 60, 80, 20, 30, 40, 70}

	bst := sortedArrayToBST(array)

	if bst.value != 50 {
		t.Errorf("%T root returned value %d, but expected %d", bst, bst.value, 50)
	}
}

// func TestTest(t *testing.T) {
// 	bst := &BinarySearchTree{5, nil, nil}

// 	values := []int{3, 7, 2, 4, 6, 8}

// 	fmt.Println(values)

// 	for _, value := range values {
// 		bst.Insert(value)
// 	}

// 	bst.Print(0)
// 	fmt.Println("--------------------")
// 	bst.Remove(5)
// 	bst.Remove(7)
// 	bst.Remove(8)
// 	bst.Print(0)
// }

// ############################# AVL TESTS #############################

func TestInsertAVL(t *testing.T) {
	bst := &BinarySearchTree{5, nil, nil}

	values := []int{3, 7, 2, 4, 6, 8}

	for index, value := range values {
		bst.Insert(value)
		size := bst.Size()
		if size != index+2 {
			t.Errorf("%T size is %d, but we expected it to be %d", bst, size, index+2)
		}
	}
}

func TestSearchAVL(t *testing.T) {
	avl := &AVL{value: 5}

	values := []int{3, 7, 2, 4, 6, 8}

	for _, value := range values {
		avl.Insert(value)

		if !avl.Search(value) {
			t.Errorf("%T does not have the value %d in its elements.", avl, value)
		}
	}

	if avl.Search(48521) {
		t.Errorf("%T does not have the value %d in its elements but said it does", avl, 48521)
	}
}

func TestRemoveAVL(t *testing.T) {
	avl := &AVL{value: 5}

	values := []int{3, 7, 2, 4, 6, 8}

	for _, value := range values {
		avl.Insert(value)
	}

	for index := len(values) - 2; index >= 0; index-- {
		avl.Remove(values[index])
		size := avl.Size()
		if size != index+2 {
			t.Errorf("%T size is %d, but we expected it to be %d", avl, size, index+2)
		}
	}
}

func TestMaxAVL(t *testing.T) {
	avl := &AVL{value: 5}

	values := []int{3, 7, 2, 4, 6, 8}

	for _, value := range values {
		avl.Insert(value)
	}

	if avl.Max() != 8 {
		t.Errorf("%T has the greatest element %d but returned %d", avl, 8, avl.Max())
	}
}

func TestMinAVL(t *testing.T) {
	avl := &AVL{value: 5}

	values := []int{3, 7, 2, 4, 6, 8}

	for _, value := range values {
		avl.Insert(value)
	}

	if avl.Min() != 2 {
		t.Errorf("%T has the smallest element %d but returned %d", avl, 2, avl.Min())
	}
}

func TestSortedArrayToAVL(t *testing.T) {
	array := []int{50, 60, 80, 20, 30, 40, 70}

	bst := sortedArrayToBST(array)

	if bst.value != 50 {
		t.Errorf("%T root returned value %d, but expected %d", bst, bst.value, 50)
	}
}

func TestRotateRightAVL(t *testing.T) {
	avl := &AVL{value: 3}

	values := []int{1, 2, 4, 5}

	for _, value := range values {
		avl.Insert(value)
	}

	if avl.value != 3 {
		t.Errorf("%T root returned value %d, but expected %d", avl, avl.value, 50)
	}

	avl = avl.RotateRight()

	if avl.value != 1 {
		t.Errorf("%T root returned value %d, but expected %d", avl, avl.value, 50)
	}
}

func TestRotateLeftAVL(t *testing.T) {
	avl := &AVL{value: 1}

	values := []int{2, 3, 4, 5}

	for _, value := range values {
		avl.Insert(value)
	}

	if avl.value != 1 {
		t.Errorf("%T root returned value %d, but expected %d", avl, avl.value, 50)
	}

	avl = avl.RotateLeft()

	if avl.value != 2 {
		t.Errorf("%T root returned value %d, but expected %d", avl, avl.value, 50)
	}
}

func TestUpdatePropertiesAVL(t *testing.T) {
	avl := &AVL{value: 1}

	values := []int{2, 3, 4, 5}

	for _, value := range values {
		avl.Insert(value)
	}

	if avl.height != 4 {
		t.Errorf("%T root returned height equals %d, but expected %d", avl, avl.height, 4)
	}

	if avl.bf != 3 {
		t.Errorf("%T root returned balance factor equals %d, but expected %d", avl, avl.bf, 3)
	}

	avl = avl.RotateLeft()

	if avl.height != 3 {
		t.Errorf("%T root returned height equals %d, but expected %d", avl, avl.value, 3)
	}

	if avl.bf != 2 {
		t.Errorf("%T root returned balance factor equals %d, but expected %d", avl, avl.value, 2)
	}
}

// func TestAVL(t *testing.T) {
// 	avl := &AVL{value: 1}

// 	values := []int{2, 3, 4, 5}

// 	fmt.Println(values)

// 	avl.Print(0)
// 	fmt.Println("Inserido 1 -> Altura(raíz):", avl.height, "BF(raíz):", avl.bf)

// 	for _, value := range values {
// 		avl.Insert(value)
// 		avl.Print(0)
// 		fmt.Println("Inserido", value, " -> Altura(raíz):", avl.height, "BF(raíz):", avl.bf)
// 		fmt.Println("--------------------")
// 	}

// 	avl.Print(0)
// 	fmt.Println("--------------------")
// 	avl = avl.RotateLeft()
// 	avl.Print(0)
// 	fmt.Println("Rotacionado à esquerda -> Altura(raíz):", avl.height, "BF(raíz):", avl.bf)
// 	fmt.Println("--------------------")
// 	avl = avl.RotateRight()
// 	avl.Print(0)
// 	fmt.Println("Rotacionado à direita (volta ao estado anterior) -> Altura(raíz):", avl.height, "BF(raíz):", avl.bf)
// 	fmt.Println("--------------------")
// 	avl = avl.RotateLeft()
// 	avl.Print(0)
// 	fmt.Println("Rotacionado à esquerda -> Altura(raíz):", avl.height, "BF(raíz):", avl.bf)
// 	fmt.Println("--------------------")
// 	avl = avl.RotateLeft()
// 	avl.Print(0)
// 	fmt.Println("Rotacionado à esquerda -> Altura(raíz):", avl.height, "BF(raíz):", avl.bf)
// }
