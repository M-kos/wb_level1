package main

import "fmt"

type Animal interface {
	castVote()
}

type Cat struct{}

func (cat *Cat) meow() {
	fmt.Println("meow")
}

type CatAdapter struct {
	cat *Cat
}

func (cat *CatAdapter) castVote() {
	cat.cat.meow()
}

func main() {
	cat := &Cat{}
	catAdapter := &CatAdapter{
		cat: cat,
	}

	Vote(catAdapter)
}

func Vote(animal Animal) {
	animal.castVote()
}

/*
Применимость паттерна описана в самом задании: паттерн Adapter позволяет сконвертировать интерфейс одного класса в интерфейс другого
Собственно плюсы вытекают из предыдущей строчки: упрощает интеграцию
Минусы: увеличивает количество структур, ухудшает чтение и поддержку
Примеры: работа с различными Апи (rest, grpc), работа с различными БДб интеграция со сторонними библиотеками, особенно когда приходится заменять одну на другую
*/
