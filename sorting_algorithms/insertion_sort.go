package main

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
