package main

/*
* Best case: Ω(nlog(n))
* Worst case: O(n²) if pivot isn't randomly taken, otherwise, O(nlog(n))
* Goal: on each iteration, choose a pivot and put elements greater than it
* at its right and lower ones at its left
**/

import (
	"math/rand"
)

func QuickSort(rand bool, list []int, args ...int) { // args is left and right elements
	if len(args) == 0 {
		args = []int{0, len(list) - 1}
	}

	left, right := args[0], args[1]
	if left < right {
		pivot := 0
		if rand {
			pivot = partitionRand(list, left, right) // if necessary randomly generated pivot
		} else {
			pivot = partition(list, left, right) // by default the pivot is the last element in the list
		}

		QuickSort(rand, list, left, pivot-1)
		QuickSort(rand, list, pivot+1, right)
	}
}

func partitionRand(list []int, left int, right int) int {
	randPivot := rand.Intn(right-left+1) + left

	list[randPivot], list[right] = list[right], list[randPivot]

	return partition(list, left, right)
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
