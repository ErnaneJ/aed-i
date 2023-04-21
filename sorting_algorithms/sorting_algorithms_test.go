package main

import (
	"reflect"
	"runtime"
	"testing"
)

var expectedResult = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var sortingFunctions = []func([]int){
	BubbleSort,
	SelectionSort,
	InsertionSort,
	MergeSort,
}

func TestAllSortingFunctions(t *testing.T) {
	for _, sortingFunction := range sortingFunctions {
		list := []int{9, 5, 1, 4, 0, 7, 8, 6, 3, 2}

		sortingFunction(list)
		currentFunctionName := runtime.FuncForPC(reflect.ValueOf(sortingFunction).Pointer()).Name()

		for i := 0; i < len(list); i++ {
			if list[i] != i {
				t.Errorf("%s: returned %d, but we expected it to be %d", currentFunctionName, list, expectedResult)
			}
		}
	}
}

func TestOnlyQuickSort(t *testing.T) {
	list := []int{9, 5, 1, 4, 0, 7, 8, 6, 3, 2}

	QuickSort(list) // This function's assignment is QuickSort(list []int, args ...int) so in this case args is [].

	for i := 0; i < len(list); i++ {
		if list[i] != i {
			t.Errorf("QuickSort: returned %d, but we expected it to be %d", list, expectedResult)
		}
	}
}
