package main

import (
	"fmt"
	"math"
)

func main() {
	var target int64
	fmt.Print("enter a number: \n")

	_, err := fmt.Scanln(&target)
	if err != nil {
		fmt.Println(err)
		return
	}

	var bitIndex int
	fmt.Print("enter the sequence number of the bit (non-negative integers): \n")

	_, err = fmt.Scanln(&bitIndex)
	if err != nil {
		fmt.Println(err)
		return
	}

	pow := bitIndex - 1
	mask := math.Pow(2, float64(pow))

	fmt.Println("result: ", target^int64(mask))
}
