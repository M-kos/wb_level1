package main

import "fmt"

func main() {
	fmt.Println(checkUniq("abcd"))
	fmt.Println(checkUniq("abCdefAaf"))
	fmt.Println(checkUniq("aabcd"))
	fmt.Println(checkUniq("атро"))
	fmt.Println(checkUniq("атроа"))
}

func checkUniq(str string) bool {
	temp := make(map[rune]struct{})
	for _, v := range str {
		if _, ok := temp[v]; ok {
			return false
		}
		temp[v] = struct{}{}
	}
	return true
}
