package main

func CountingSort(list []int) {
	smaller := list[0]
	bigger := list[0]
	listLength := len(list)

	for i := 1; i < listLength; i++ {
		if list[i] < smaller {
			smaller = list[i]
		}

		if list[i] > bigger {
			bigger = list[i]
		}
	}

	rangeList := bigger - smaller + 1

	count := make([]int, rangeList)

	for i := 0; i < listLength; i++ {
		count[list[i]-smaller]++
	}

	for i := 1; i < rangeList; i++ {
		count[i] += count[i-1]
	}

	copyList := make([]int, listLength)
	copy(copyList, list)

	busy := make([]bool, len(copyList))

	for i := 0; i < listLength; i++ {
		sortedIndex := count[copyList[i]-smaller] - 1
		for busy[sortedIndex] {
			sortedIndex--
		}
		list[sortedIndex] = copyList[i]
		busy[sortedIndex] = true
	}
}
