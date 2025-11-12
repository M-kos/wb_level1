package main

import (
	"fmt"
	"github.com/M-kos/wb_level1/task_22/internal/calc"
	"math/big"
)

func main() {
	xInt, yInt := 10, 3
	fmt.Println(calc.Add(xInt, yInt), calc.Sub(xInt, yInt), calc.Multiple(xInt, yInt), calc.Div(xInt, yInt))

	var xInt64 int64 = 500
	var yInt64 int64 = 230
	fmt.Println(calc.Add(xInt64, yInt64), calc.Sub(xInt64, yInt64), calc.Multiple(xInt64, yInt64), calc.Div(xInt64, yInt64))

	xFloat, yFloat := 3.14, 2.0
	fmt.Println(calc.Add(xFloat, yFloat), calc.Sub(xFloat, yFloat), calc.Multiple(xFloat, yFloat), calc.Div(xFloat, yFloat))

	bigX := big.NewInt(10000000)
	bigY := big.NewInt(50000000)

	fmt.Println(calc.Add(bigX, bigY), calc.Sub(bigX, bigY), calc.Multiple(bigX, bigY), calc.Div(bigX, bigY))

	bigFloatX := big.NewFloat(123213.12)
	bigFloatY := big.NewFloat(98765.4321)

	fmt.Println(calc.Add(bigFloatX, bigFloatY), calc.Sub(bigFloatX, bigFloatY), calc.Multiple(bigFloatX, bigFloatY), calc.Div(bigFloatX, bigFloatY))
}
