package main

func recursiveBinarySearch(list []int, start int, end int, element int) int {
	defer timer("Recursive Binary Search")()
	if end >= len(list) {
		end = len(list) - 1
	}
	if start > end {
		return -1
	}

	mid := (start + end) / 2
	if element == list[mid] {
		return mid
	} else if element < list[mid] {
		return recursiveBinarySearch(list, start, mid-1, element)
	} else {
		return recursiveBinarySearch(list, mid+1, end, element)
	}
}
