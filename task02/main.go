package main

import (
	"fmt"
	"sync"
)

// Функция, которая выводит число и его квадрат на терминал
func Square(a int, wg *sync.WaitGroup) {
	fmt.Printf("%d squared equals %d\n", a, a * a)
	// Уменьшает число ожидаемых горутин
	wg.Done()
}

func main() {
	// Инициализируем массив 
	arr := [5]int{2, 4, 6, 8, 10}

	// Инициализируем WaitGroup для того, чтобы главная горутина ожидала выполнения остальных
	var wg sync.WaitGroup

	for _, value := range arr {
		// Увеличивает число ожидаемых горутин
		wg.Add(1)
		// Запускаем горутину
		go Square(value, &wg)
	}


	// Ожидаем завершения всех горутин
	wg.Wait()

}