package main

import "fmt"


// Функция для удаление элемента по индексу
func Remove(slice []int, targetIndex int) []int {
	length := len(slice)
	if targetIndex < 0 || targetIndex >= length {
		return slice
	}
	slice = append(slice[:targetIndex], slice[targetIndex+1:]...)
	return slice
}



func main(){
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Array before removal: %v\n", slice)

	slice = Remove(slice, 0)
	
	fmt.Printf("Array after removal: %v\n", slice)
}