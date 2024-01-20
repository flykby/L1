package main

import (
	"fmt"
	"time"
)


// Отправляет данные в первый канал
func Sender(ch1 chan int, arr []int) {
	for _, v := range arr {
		ch1 <- v
		time.Sleep(time.Second)
	}
	close(ch1)
}


// Модифицирует и переадресует данные в канал 2
func Redirector(ch1 chan int, ch2 chan int) {
	for {
		n, ok := <- ch1
		if !ok {
			fmt.Printf("Redirector is closed\n")
			return
		}
		ch2 <- n * n
	}
}


// Печатает данные из канала 2
func Writer(ch2 chan int) {
	for {
		n, ok := <- ch2
		if !ok {
			fmt.Println("Writer is closed")
		}
		fmt.Printf("Get data: %d\n", n)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	arr := []int{1, 5, 3, 10, 2, 6, 8, 11}
	

	go Writer(ch2)
	go Redirector(ch1, ch2)
	Sender(ch1, arr)

}