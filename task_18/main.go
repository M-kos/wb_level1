package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
)

type Counter struct {
	count atomic.Int64
}

func (c *Counter) Increment() {
	c.count.Add(1)
}

func (c *Counter) Decrement() {
	c.count.Add(-1)
}

func (c *Counter) Value() int64 {
	return c.count.Load()
}

func (c *Counter) Reset() {
	c.count.Store(0)
}

func main() {
	counter := &Counter{}
	wg := sync.WaitGroup{}

	for range 10 {
		wg.Add(1)
		k := rand.Intn(6)

		if k > 3 {
			fmt.Println("inc")
			go func() {
				defer wg.Done()
				counter.Increment()
			}()
		}

		if k <= 3 {
			fmt.Println("dec")
			go func() {
				defer wg.Done()
				counter.Decrement()
			}()
		}
	}

	wg.Wait()
	fmt.Println("value", counter.Value())
}
