package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)



func main() {
	// Создаем канал для записи и чтения данных
	ch := make(chan int)
	// С помощью flag определяем кол-во воркеров, по-дефолту 4
	n := flag.Int("n", 4, "workers")
	flag.Parse()
	// Создаем wg для ожидания главной горутиной остальных
	wg := sync.WaitGroup{}
	wg.Add(*n)
	
	// Регистрируем обработчик сигнала прерывания 
	notifyContext, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Запускаем n воркеров, каждый в отдельной горутине
	for i := 0; i < *n; i++ {
		go Worker(&wg, ch, i)
	}

	// Запускаем горутину отсылающую в канал данные
	go func() {
		for {
			select {
			// Отсылаем данные
			case ch <- rand.Intn(900) + 100:
				time.Sleep(time.Millisecond * 50)
			case <- notifyContext.Done(): // Ожидаем получения сигнала завершения
				close(ch)
				return
			}
		}
	}()

	// Ожидаем завершения запущенных горутин
	wg.Wait()

}

// Воркер считывает данные из канала и выводит их в терминал
func Worker(wg *sync.WaitGroup, ch chan int, id int) {
	defer wg.Done()

	for {
		data, ok := <-ch
		if !ok {
			fmt.Printf("Worker id: %d closed\n", id)
			return
		}
		fmt.Println("Worker id:", id, "get data:", data)
		time.Sleep(time.Millisecond * 100)
	}
}