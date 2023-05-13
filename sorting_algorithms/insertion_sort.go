package main

/*
* Best case: Ω(n)
* Worst case: O(n²)
* Goal: insert the i-th element at its correct position on the "left hand"
**/

func InsertionSort(list []int) {
	size := len(list)
	for i := 1; i < size; i++ {
		performedTheExchange := false
		for j := size - 1; j > i-1; j-- {
			if list[j] < list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
				performedTheExchange = true
			}
		}
		if !performedTheExchange {
			return // it's already sorted
		}
	}
}
