package main

func interactiveBinarySearch(list []int, element int) int {
	defer timer("Interactive Binary Search")()
	start := 0
	end := len(list)

	for start < end {
		mid := (start + end) / 2
		if element == list[mid] {
			return mid
		} else if element < list[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}
