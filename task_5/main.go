package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	inChan := make(chan int)
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		wrighter(ctx, inChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		reader(ctx, inChan)

	}()

	wg.Wait()
	fmt.Println("all workers done")
	<-ctx.Done()
}

func wrighter(ctx context.Context, outChan chan<- int) {
	defer close(outChan)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("wrighter done")
			return
		case outChan <- rand.Int():
		}
	}
}

func reader(ctx context.Context, inChan <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("reader done")
			return
		case v := <-inChan:
			fmt.Println(v)
		}
	}
}
