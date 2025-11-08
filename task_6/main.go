package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	flagCondition()
	doneChanel()
	canceledContext()
	contextWithTimeout()
	runtimeGoExit()
	osExit()
	isPanic()

	fmt.Println("finish")
}

func flagCondition() {
	done := false
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if done {
				fmt.Println("[flag] goroutine stopped")
				return
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()

	time.Sleep(time.Second * 2)
	done = true
	wg.Wait()
}

func doneChanel() {
	done := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				fmt.Println("[done ch] goroutine stopped")
				return
			default:
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	close(done)
	wg.Wait()
}

func canceledContext() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("[context with cancel] goroutine stopped")
				return
			default:
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	cancel()
	wg.Wait()
}

func contextWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("[context with timeout] goroutine stopped")
				return
			default:
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()
}

func runtimeGoExit() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("[runtime exit] goroutine stopped")
		time.Sleep(time.Second * 1)
		runtime.Goexit()
	}()

	wg.Wait()
}

func osExit() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 1)
		fmt.Println("[os exit] goroutine will stop")
		os.Exit(1)
	}()

	wg.Wait()
}

func isPanic() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("[panic] goroutine stopped")
			}
		}()
		time.Sleep(time.Second * 1)
		panic("stop it")
	}()

	time.Sleep(2 * time.Second)
}
