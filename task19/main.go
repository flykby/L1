package main

import (
	"fmt"
)

func main() {
	str := "Bad != good"
	fmt.Println(str)
	// Создаем массив символов из строки
	runes := []rune(str)
	lastIndex := len(runes) - 1
	// Переворачиваем строку
	for i := 0; i < len(runes) / 2; i++ {
		runes[i], runes[lastIndex - i] = runes[lastIndex - i], runes[i]
	}
	// Конвертируем обратно массив символов в строку
	str = string(runes)
	fmt.Println(str)
}