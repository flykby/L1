package main

import "fmt"


func Merge(arr1 []int, arr2 []int) []int {
	set := make(map[int]interface{})
	for _, v := range arr1 {
		set[v] = nil
	}
	for _, v := range arr2 {
		_, ok := set[v]
		if !ok {
			set[v] = nil
			arr1 = append(arr1, v)
		}
	}
	return arr1
}

func main() {
	arr1 := []int {3, 2, 6, 9, 13, 18}
	arr2 := []int {1, 2, 12, 7, 3}

	fmt.Println(Merge(arr1, arr2))
}