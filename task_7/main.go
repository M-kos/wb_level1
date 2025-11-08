package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	m := make(map[int]int)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			key := rand.Intn(3)
			value := rand.Intn(10)

			fmt.Printf("key: %d, value: %d\n", key, value)

			mu.Lock()
			m[key] = value
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("finish: ", m)
}
