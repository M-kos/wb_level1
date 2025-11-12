package main

import "fmt"

func main() {
	v1()
	v1()
}

func v1() {
	a := []int{1, 2, 3, 4, 5, 6}
	i := 2

	fmt.Println(a, len(a), cap(a))

	copy(a[i:], a[i+1:])
	a = a[:len(a)-1]

	fmt.Println(a, len(a), cap(a))
}

func v2() {
	a := []int{1, 2, 3, 4, 5, 6}
	i := 2

	fmt.Println(a, len(a), cap(a))

	var temp []int
	a = append(temp, a[:i]...)
	a = append(a, a[i+1:]...)

	fmt.Println(a, len(a), cap(a))
}
