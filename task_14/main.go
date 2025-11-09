package main

import "fmt"

func main() {
	a := 5
	b := "str"
	c := false
	d := make(chan interface{})

	checkType(a)
	checkType(b)
	checkType(c)
	checkType(d)
}

func checkType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan interface{}:
		fmt.Println("chan")
	}
}
