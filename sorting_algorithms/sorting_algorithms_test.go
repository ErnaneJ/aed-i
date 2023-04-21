package main

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	list := []int{9, 5, 1, 4, 0, 7, 8, 6, 3, 2}

	BubbleSort(list)

	if list[0] != 0 || list[len(list)-1] != 9 {
		t.Errorf("BubbleSort: returned %d, but we expected it to be %d", list, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	}
}

func TestSelectionSort(t *testing.T) {
	list := []int{9, 5, 1, 4, 0, 7, 8, 6, 3, 2}

	SelectionSort(list)

	if list[0] != 0 || list[len(list)-1] != 9 {
		t.Errorf("SelectionSort: returned %d, but we expected it to be %d", list, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	}
}

func TestInsertionSort(t *testing.T) {
	list := []int{9, 5, 1, 4, 0, 7, 8, 6, 3, 2}

	InsertionSort(list)

	if list[0] != 0 || list[len(list)-1] != 9 {
		t.Errorf("SelectionSort: returned %d, but we expected it to be %d", list, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	}
}

func TestMergeSort(t *testing.T) {
	list := []int{0, 2, 1, 3, 4, 5, 6, 7, 8, 9}

	MergeSort(list)

	if list[0] != 0 || list[len(list)-1] != 9 {
		t.Errorf("MergeSort: returned %d, but we expected it to be %d", list, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	}
}
