package main

import "fmt"

// Базовый алгоритм бинарного поиска
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr) - 1

	for left <= right {
		middle := ((right - left) >> 1) + left

		if arr[middle] == target {
			return middle
		}
		if arr[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}

func main() {
	arr := []int{0, 4, 14, 15, 20, 33, 34, 87, 100}
	fmt.Println(binarySearch(arr, 112))
}