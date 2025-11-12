package main

import (
	"fmt"
	"github.com/M-kos/wb_level1/task_25/internal/timer"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	timer.Sleep(time.Second * 3)
	fmt.Println(time.Since(now))
}
