package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "snow dog sun"
	b := "привет мир"
	c := ""
	d := "просто строка just string"

	fmt.Println(revertWord(a))
	fmt.Println(revertWord(b))
	fmt.Println(revertWord(c))
	fmt.Println(revertWord(d))
}

func revertWord(s string) string {
	var result strings.Builder
	lastIdx := len(s)

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 32 {
			result.WriteString(s[i+1 : lastIdx])
			result.WriteString(" ")
			lastIdx = i
		}

		if i == 0 {
			result.WriteString(s[i:lastIdx])
		}
	}

	return result.String()
}
