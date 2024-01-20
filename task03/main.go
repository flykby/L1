package main

import (
	"fmt"
	"sync"
)

// Функция, которая прибавляет квадрат числа к сумме
func Sum(a int, sum *int, wg *sync.WaitGroup) {
	// Прибавляем квадрат к сумме чисел
	*sum += a * a
	// Уменьшает число ожидаемых горутин
	wg.Done()
}

func main() {
	// Инициализируем массив 
	arr := [5]int{2, 4, 6, 8, 10}

	// Инициализируем WaitGroup для того, чтобы главная горутина ожидала выполнения остальных
	var wg sync.WaitGroup

	sum := 0

	for _, value := range arr {
		// Увеличивает число ожидаемых горутин
		wg.Add(1)
		// Запускаем горутину
		go Sum(value, &sum, &wg)
	}

	
	// Ожидаем завершения всех горутин
	wg.Wait()
	// Выводим результат
	fmt.Printf("Sum: %d\n", sum)
}