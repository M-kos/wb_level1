package main

import "fmt"

func main() {
	variant1()
	variant2()
}

func variant1() {
	a := 5
	b := 1

	fmt.Println("a: ", a, " b: ", b)

	a, b = b, a

	fmt.Println("a: ", a, " b: ", b)
}

func variant2() {
	a := 5
	b := 1

	fmt.Println("a: ", a, " b: ", b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println("a: ", a, " b: ", b)
}
