package main

import (
	"fmt"
	"github.com/M-kos/wb_level1/task_24/internal/point"
)

func main() {
	a := point.NewPoint(3, 2)
	b := point.NewPoint(7, 8)

	fmt.Println(a.DistanceTo(b))
}
