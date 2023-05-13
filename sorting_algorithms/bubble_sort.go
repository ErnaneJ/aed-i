package main

/*
* Best case: Ω(n)
* Worst case: O(n²)
* Goal: move the highest element to the right
**/

func BubbleSort(list []int) {
	size := len(list)
	for i := 0; i < size-1; i++ {
		performedTheExchange := false
		for j := 0; j < size-i-1; j++ {
			if list[j] > list[j+1] {
				performedTheExchange = true
				list[j+1], list[j] = list[j], list[j+1]
			}
		}
		if !performedTheExchange {
			return // it's already sorted
		}
	}
}
