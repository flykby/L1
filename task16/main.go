package main

import "fmt"

// Алгоритм быстрой сортировки

func quicksort(arr []int, left int, right int) {
    if left < right {
        part := partition(arr, left, right)
        quicksort(arr, left, part-1)
        quicksort(arr, part+1, right)
    }
}

func partition(arr []int, left int, right int) int {
    pivot := arr[right]
    i := left - 1
    for j := left; j <= right-1; j++ {
        if arr[j] <= pivot {
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    arr[i+1], arr[right] = arr[right], arr[i+1]
    return i+1
}

func main() {
	arr := []int{1, 5, 7, 3, 9, 2, 4, 6, 8, 0}
	fmt.Println("Unsorted array:", arr)
	quicksort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:", arr)
}