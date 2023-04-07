package main

import "testing"

func TestFound(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	element := 5

	resultIBS := interactiveBinarySearch(list, element)
	resultRBS := recursiveBinarySearch(list, 0, len(list), element)
	resultSS := sequentialSearch(list, element)

	if resultIBS != 5 {
		t.Errorf("interactiveBinarySearch: returned %d, but we expected it to be %d", resultIBS, 5)
	}
	if resultRBS != 5 {
		t.Errorf("recursiveBinarySearch: returned %d, but we expected it to be %d", resultRBS, 5)
	}
	if resultSS != 5 {
		t.Errorf("sequentialSearch: returned %d, but we expected it to be %d", resultSS, 5)
	}
}

func TestNotFound(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	element := 50

	resultIBS := interactiveBinarySearch(list, element)
	resultRBS := recursiveBinarySearch(list, 0, len(list), element)
	resultSS := sequentialSearch(list, element)

	if resultIBS != -1 {
		t.Errorf("interactiveBinarySearch: returned %d, but we expected it to be %d", resultIBS, -1)
	}
	if resultRBS != -1 {
		t.Errorf("recursiveBinarySearch: returned %d, but we expected it to be %d", resultRBS, -1)
	}
	if resultSS != -1 {
		t.Errorf("sequentialSearch: returned %d, but we expected it to be %d", resultSS, -1)
	}
}
