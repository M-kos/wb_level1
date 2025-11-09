package main

import "fmt"

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 0, -1, 1, -102}
	result := make(map[int][]float64)
	step := 10.0

	for _, temperature := range temps {
		k := int(temperature / step)
		group := k * int(step)

		result[group] = append(result[group], temperature)
	}

	fmt.Println(result)
}
