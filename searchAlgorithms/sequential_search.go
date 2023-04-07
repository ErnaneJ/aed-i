package main

func sequentialSearch(list []int, item int) int {
	defer timer("Sequencial Search")()
	for i := 0; i < len(list); i++ {
		if list[i] == item {
			return i
		}
	}
	return -1
}
