package main

/*
* Best case: Ω(nlog(n))
* Worst case: O(nlog(n))
* Goal: split array into halves with recursion calls. Then,
* given two halves that are sorted, merge them into a sorted array
**/

func MergeSort(list []int) {
	if len(list) <= 1 {
		return
	}

	mid := len(list) / 2

	left := make([]int, mid)
	right := make([]int, len(list[mid:]))

	copy(left, list[:mid])
	copy(right, list[mid:])

	MergeSort(left)
	MergeSort(right)

	merge(list, left, right)
}

func merge(list []int, left []int, right []int) {
	lenLeft, lenRight := len(left), len(right)
	i, j, k := 0, 0, 0

	for i < lenLeft && j < lenRight {
		if left[i] < right[j] {
			list[k] = left[i]
			i++
		} else {
			list[k] = right[j]
			j++
		}
		k++
	}

	for i < lenLeft {
		list[k] = left[i]
		k++
		i++
	}

	for j < lenRight {
		list[k] = right[j]
		k++
		j++
	}
}
