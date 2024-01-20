package main

import (
	"fmt"
	"task24/point"
)


func main() {
	// Инициализируем 2 точки
	p1 := *point.NewPoint(0.0, 4.0)
	p2 := *point.NewPoint(4.0, 1.0)

	// Выводим работу метода, вычисляющего расстоянию между двумя точками
	fmt.Println(p1.GetDistance(&p2))
}
