package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
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
	var wg sync.WaitGroup
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, inChan)
		}()
	}

	go func() {
		wg.Add(1)
		defer wg.Done()
		defer func() {
			close(inChan)
		}()

		for i := range 1000000000 {
			select {
			case <-ctx.Done():
				return
			case inChan <- i:
			}
		}

	}()

	wg.Wait()
	fmt.Println("all workers done")
	<-ctx.Done()
}

func worker(ctx context.Context, inChan <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker done")
			return
		case v := <-inChan:
			fmt.Println(v)
		}
	}
}
