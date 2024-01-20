package main

import (
	"fmt"
	"math/rand"
	"time"
)

func read(ch chan int) {
	for {
		// Получаем данные и выводим их
		fmt.Println("Get data:", <- ch)
		
		time.Sleep(time.Millisecond * 500)
	}
}

func write(ch chan int) {
	for {
		// Генерируем данные
		a := rand.Intn(900) + 100
		// Отправляем данные в канал
		ch <- a
		// Выводим в терминал отправленные данные
		fmt.Println("Send data:", a)
		// Слипаем для удобства в отслеживании логов
		time.Sleep(time.Microsecond * 500)
	}
}

func main() {
	ch := make(chan int)

	go write(ch)
	go read(ch)

	time.Sleep(time.Second * 5)
}