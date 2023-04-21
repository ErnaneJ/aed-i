package main

func QuickSort(list []int, args ...int) { // args is left and right elements
	if len(args) == 0 {
		args = []int{0, len(list) - 1}
	}

	left, right := args[0], args[1]
	if left < right {
		pivot := partition(list, left, right)
		QuickSort(list, left, pivot-1)
		QuickSort(list, pivot+1, right)
	}
}

func partition(list []int, left, right int) int {
	pivot := list[right]
	smallersIterator := left - 1

	for largestIterator := left; largestIterator < right; largestIterator++ {
		if list[largestIterator] < pivot {
			smallersIterator++
			list[smallersIterator], list[largestIterator] = list[largestIterator], list[smallersIterator]
		}
	}

	list[smallersIterator+1], list[right] = list[right], list[smallersIterator+1]
	return smallersIterator + 1
}
