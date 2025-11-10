package main

import "fmt"

func main() {
	a := "главрыба"
	b := "робот"
	c := "hello"
	d := ""
	e := "Привет, Golang"
	f := "🚀 🎒 🎉"

	fmt.Println(revertStr(a))
	fmt.Println(revertStr(b))
	fmt.Println(revertStr(c))
	fmt.Println(revertStr(d))
	fmt.Println(revertStr(e))
	fmt.Println(revertStr(f))
}

func revertStr(in string) string {
	runeStr := []rune(in)
	result := make([]rune, 0, len(runeStr))

	for i := len(runeStr) - 1; i >= 0; i-- {
		result = append(result, runeStr[i])
	}

	return string(result)
}
