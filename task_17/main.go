package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{-4, -1, 0, 1, 5, 7, 8, 10, 12}
	a1 := []int{-4, -1, 0, 1, 5, 7, 8, 10, 12}
	b := make([]int, 0)
	c := []int{-4, 5}
	d := []int{10}
	fmt.Println(binarySearch(a, -1))
	fmt.Println(binarySearch(a1, 25))
	fmt.Println(binarySearch(b, 8))
	fmt.Println(binarySearch(c, 5))
	fmt.Println(binarySearch(d, 11))
}

func binarySearch(in []int, target int) int {
	pivotIdx := getPivotIdx(in)

	for {
		if len(in) == 0 {
			break
		}

		if in[pivotIdx] == target {
			return pivotIdx
		}

		if target > in[pivotIdx] {
			in = in[pivotIdx+1:]
		} else {
			in = in[:pivotIdx]
		}

		pivotIdx = getPivotIdx(in)
	}

	return -1
}

func getPivotIdx(in []int) int {
	return int(math.Floor(float64(len(in)) / 2))
}
