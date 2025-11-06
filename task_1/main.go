package main

import "fmt"

type Human struct {
}

func (h Human) SayHello() {
	fmt.Println("Hello")
}

type Action struct {
	Human
}

func main() {
	var act Action

	act.SayHello()
}
