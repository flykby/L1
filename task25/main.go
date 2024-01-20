package main

import (
	"context"
	"fmt"
	"time"
)

// Ожидаем выполнение функции time.After(duration) и завершает работу
func Sleep1(duration time.Duration) {
	<-time.After(duration)
}

// Создаем таймаут, имеющий длительность в duration и ожидаем его завершения.
// После отменяем контекст через cancelFunc()
func Sleep2(duration time.Duration) {
	timeout, cancelFunc := context.WithTimeout(context.Background(), duration)
	defer cancelFunc()
	<-timeout.Done()
}


func main() {
	fmt.Println("Ожидаем 5 секунд")
	Sleep1(3 * time.Second)
	fmt.Println("Время истекло")

	fmt.Println("Ожидаем 5 секунд")
	Sleep2(3 * time.Second)
	fmt.Println("Время истекло")

}