package main

import "fmt"

// Проблема в не эффективном использовании памяти:
// Если мы используем лишь первые 100 элементов
// из выделенных 1024, то 924 элементов не будут удалены,
// и будут висеть в памяти, сборщик мусора удалит это лишь
// тогда - когда justString тоже выйдет из области видимости



var justString *string

func someFunc() *string {
    v1 := []rune(createHugeString(1 << 10))
	v2 := make([]rune, 100)

	copy(v2, v1[:100]) // Копируем данные в новую область памяти

	justString := string(v2)
    return &justString
}

func main() {
    str := *someFunc() // Получаем указатель на строку
	fmt.Println(str)
}

