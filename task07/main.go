package main

import (
	"fmt"
	"sync"
)

// Создаем структуру с мапой и mutex-ом
type Data struct {
	data map[int]int
	mutex sync.Mutex
}


func main() {
	n := 100
	// Создаем экзмепляр структуры 
	data := &Data{make(map[int]int), sync.Mutex{}}

	// Создаем wg для ожидания выполнения горутин
	wg := sync.WaitGroup{}
	wg.Add(n)

	// В цикле запускаем n горутин, которые конкурентно пишут данные в мапу
	for i := 0; i < n; i++ {
		go func(v int) {
			// Блокируем mutex
			data.mutex.Lock()
			// Пишем данные в мапу
			data.data[v] = v
			// Разблокируем mutex
			data.mutex.Unlock()
			// Говорим wg, что 1 поток завершен
			wg.Done()
		}(i)
	}
	// Ожидаем завершения горутин
	wg.Wait()

	for k, v := range data.data {
		fmt.Printf("Key: %d. Value: %d\n", k, v)
	}

}