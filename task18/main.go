package main

import (
	"fmt"
	"sync"
)


// Структура для хранения счетчика и mutex для синхронизации потоков
type Incrementor struct {
	data int
	mutex sync.Mutex
}

// Метод для увеличения счетчика
func (i *Incrementor) IncreaseСounter() {
	defer i.mutex.Unlock()
	i.mutex.Lock()
	i.data += 1
}

func main() {
	// Инициализируем структуру счетчика
	incrementor := &Incrementor{data: 0}
	// Инициализируем структуру для ожидания группы горутин
	wg := sync.WaitGroup{}
	
	// Запускаем 3 горутины
	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func (inc *Incrementor) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				inc.IncreaseСounter()
			}
		}(incrementor)
	} 
	// Ожидаем завершение работы горутин
	wg.Wait()

	fmt.Println(incrementor.data)

}