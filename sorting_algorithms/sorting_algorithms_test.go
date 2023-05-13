package main

import (
	"math/rand"
	"reflect"
	"runtime"
	"testing"
	"time"
)

var expectedResult = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var sortingFunctions = []func([]int){
	BubbleSort,
	SelectionSort,
	InsertionSort,
	MergeSort,
	CountingSort,
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
	rand.Seed(time.Now().UnixNano())
	// QuickSort function's assignment is QuickSort(rand bool, list []int, args ...int)
	// so in this case args is [].

	list := []int{9, 5, 1, 4, 0, 7, 8, 6, 3, 2}

	QuickSort(true, list) // Enable random pivot

	for i := 0; i < len(list); i++ {
		if list[i] != i {
			t.Errorf("QuickSort: returned %d, but we expected it to be %d", list, expectedResult)
		}
	}

	list = []int{9, 5, 1, 4, 0, 7, 8, 6, 3, 2}

	QuickSort(false, list) // pivot is always the last element

	for i := 0; i < len(list); i++ {
		if list[i] != i {
			t.Errorf("QuickSort: returned %d, but we expected it to be %d", list, expectedResult)
		}
	}
}

func TestExamples(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	list := []int{1, 2, 3, 3, 4, 5, 6, 7}
	QuickSort(true, list)
}
