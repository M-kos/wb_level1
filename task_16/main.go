package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{10, -4, 5, 8, -1, 0, 1, 7, 12}
	b := make([]int, 0)
	c := []int{4, -5}
	d := []int{10}
	fmt.Println(quickSort(a))
	fmt.Println(quickSort(b))
	fmt.Println(quickSort(c))
	fmt.Println(quickSort(d))
}

func quickSort(in []int) []int {
	if len(in) == 0 || len(in) == 1 {
		return in
	}

	result := make([]int, 0, len(in))
	pivotIdx := int(math.Floor(float64(len(in)) / 2))

	left := make([]int, 0)
	right := make([]int, 0)

	for i, n := range in {
		if i == pivotIdx {
			continue
		}

		if n >= in[pivotIdx] {
			right = append(right, n)
		} else {
			left = append(left, n)
		}
	}

	leftResult := quickSort(left)
	rightResult := quickSort(right)

	result = append(result, leftResult...)
	result = append(result, in[pivotIdx])
	result = append(result, rightResult...)

	return result
}
