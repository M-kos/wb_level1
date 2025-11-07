package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := [...]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, v := range nums {
		wg.Add(1)

		go func() {
			defer wg.Done()
			fmt.Println(v * v)
		}()
	}

	wg.Wait()
}
