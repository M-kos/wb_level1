package main

import "fmt"

func main() {
	a := []string{"cat", "cat", "dog", "cat", "tree"}
	temp := make(map[string]struct{})

	for _, v := range a {
		temp[v] = struct{}{}
	}

	result := make([]string, 0, len(temp))

	for k := range temp {
		result = append(result, k)
	}

	fmt.Println(result)
}
