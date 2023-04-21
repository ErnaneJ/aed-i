package main

func SelectionSort(list []int) {
	size := len(list)
	for i := 0; i < size; i++ {
		minIndex := i
		for j := i + 1; j < size; j++ {
			if list[minIndex] > list[j] {
				minIndex = j
			}
		}

		list[minIndex], list[i] = list[i], list[minIndex]
	}
}
