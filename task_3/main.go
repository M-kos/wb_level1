package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

func main() {
	var workersCount int

	flag.IntVar(&workersCount, "count", 1, "Number of workers")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("workers count must be passed")
		return
	}

	inChan := make(chan int)
	wg := sync.WaitGroup{}

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range inChan {
				fmt.Println(v)
			}
		}()
	}

	for i := range 1000000 {
		inChan <- i
	}

	close(inChan)
	wg.Wait()
}
