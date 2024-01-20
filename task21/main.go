package main

import "fmt"

// Адаптер - структурный паттерн проектирования, который позволяет объектам с разными
// интерфейсами работать вместе. Он позволяет адаптировать интерфейс одного класса к 
// другому, независимо от их реализаций.

// Структуры Орел и Кошка, реализующий каждый свой интерефейс
type Bird interface {
	Fly()
}

type Animal interface {
	Walk()
}

type Eagle struct {

}

func (e *Eagle) Fly() {
	fmt.Println("I'm flying!")
}

type Cat struct {

}

func (c *Cat) Walk() {
	fmt.Println("I'm walking!")
}

// Утка(Адаптер), реализующий оба интерфейса
type Duck struct {
	bird Bird
	animal Animal
}

func BornDuck(a Animal, b Bird) *Duck {
	return &Duck {
		bird: b,
		animal: a,
	}
}


func main() {
	eagle := Eagle{}
	cat := Cat{}

	duck := BornDuck(&cat, &eagle)

	duck.animal.Walk()
	duck.bird.Fly()

}