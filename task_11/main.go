package main

import "fmt"

func main() {
	a := []int{1, 3, 2}
	b := []int{2, 4, 3}

	temp := make(map[int]struct{})
	result := make([]int, 0)

	for _, v := range a {
		temp[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := temp[v]; ok {
			delete(temp, v)
			result = append(result, v)
		}
	}

	fmt.Println(result)
}
