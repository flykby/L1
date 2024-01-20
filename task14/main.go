package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{}
	a = 42
	// Вычисляем с помощью рефлексии тип данных
	fmt.Println(reflect.TypeOf(a))
	a = "press F"
	fmt.Println(reflect.TypeOf(a))
	a = true
	fmt.Println(reflect.TypeOf(a))
	a = new(chan int)
	fmt.Println(reflect.TypeOf(a))

	
}