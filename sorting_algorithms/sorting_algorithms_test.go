package main

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	list := []int{9, 5, 1, 4, 7, 8, 6, 3, 2}

	BubbleSort(list)

	if list[0] != 1 || list[len(list)-1] != 9 {
		t.Errorf("BubbleSort: returned %d, but we expected it to be %d", list, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	}
}

func TestSelectionSort(t *testing.T) {
	list := []int{9, 5, 1, 4, 7, 8, 6, 3, 2}

	SelectionSort(list)

	if list[0] != 1 || list[len(list)-1] != 9 {
		t.Errorf("SelectionSort: returned %d, but we expected it to be %d", list, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	}
}

func TestInsertionSort(t *testing.T) {
	list := []int{9, 5, 1, 4, 7, 8, 6, 3, 2} // => 36
	// list := []int{2, 1, 3, 4, 5, 6, 7, 8, 9} // => 15

	InsertionSort(list)

	if list[0] != 1 || list[len(list)-1] != 9 {
		t.Errorf("SelectionSort: returned %d, but we expected it to be %d", list, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	}
}
