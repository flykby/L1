package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)


func routineContext(ctx context.Context, wg *sync.WaitGroup, msg string) {
	defer wg.Done()
	<-ctx.Done()
	fmt.Println("\rClosed", msg)
}

func ExampleWithCancel() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ctx, cancelFunc := context.WithCancel(context.Background())
	go routineContext(ctx, &wg, "1 - With cancel")
	time.Sleep(time.Second)
	cancelFunc()
	wg.Wait()
}

func ExampleWithTimeout() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ctxTime, cancelFuncTime := context.WithTimeout(context.Background(), time.Second)
	defer cancelFuncTime()
	go routineContext(ctxTime, &wg, "2 - With timeout")
	<-ctxTime.Done()
	wg.Wait()
}

func ExampleWithCtxSignal() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	fmt.Println("\tPress CTRL+C to close notify context goroutine")
	deadline, cancelFunc := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancelFunc()
	go routineContext(deadline, &wg, "4 with context signal CTRL+C")
	<-deadline.Done()
	wg.Wait()
}


func main() {
	// С помощью функции cancel() который получили с контекста, мы имеем возможность сообщить горутинам о прекращении работы.
	ExampleWithCancel()
	// Контекст в котором через заданное время будет послан сигнал о прекращении работы горутинам
	ExampleWithTimeout()
	// Тоже самое что и выше, но привязанный к сигналам, в данном случае (ctrl + c)
	ExampleWithCtxSignal()
}