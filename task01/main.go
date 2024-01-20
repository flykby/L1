package main

import "fmt"

// Создаем стуктуру Human
type Human struct {
	Name string
	Age  int
}

// Определяем метод SayGreet, который выводит приветствие
func (h *Human) SayGreet() {
	greet := fmt.Sprintf("Hi! My name is %s. I am %d years old.", h.Name, h.Age)
	fmt.Println(greet)
}

// Создаем структуру Action и наследуем ее встраиванием от Human
type Action struct {
	Human
}

func main() {
	// Инициализируем переменную типа Human
	human := Human {
		Name : "Mark",
		Age: 15,
	}
	// Вызываем метод SayGreet
	human.SayGreet()
	
	// Инициализируем переменную типа Action
	action := &Action{
		human,
	}
	
	// Вызываем метод SayGreet
	action.SayGreet()

	fmt.Println("")
	
	// Переопределяем поля у обьекта action
	action.Name = "Mark II"
	action.Age = 16
	
	// Вызываем метод SayGreet
	human.SayGreet()
	action.SayGreet()

}