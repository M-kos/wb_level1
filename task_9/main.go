package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	numsCh := generator(nums)
	resCh := multiplier(numsCh)

	for n := range resCh {
		fmt.Println(n)
	}
}

func generator(nums []int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, n := range nums {
			out <- n
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func multiplier(in chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for n := range in {
			out <- n * 2
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
