package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Bad bars age"
	fmt.Println(str)
	// Создаем массив символов из строки
	strSlice := strings.Split(str, " ")
	lastIndex := len(strSlice) - 1
	// Переворачиваем строку
	for i := 0; i < len(strSlice) / 2; i++ {
		strSlice[i], strSlice[lastIndex - i] = strSlice[lastIndex - i], strSlice[i]
	}
	// Конвертируем обратно массив символов в строку
	str = strings.Join(strSlice, " ")
	fmt.Println(str)
}