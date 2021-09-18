package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		work()
	}()
	// join point needed
	wg.Wait()
	fmt.Println("elapse time:", time.Since(now))
	fmt.Println("done waiting, main exits!")
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("printing some stuff")
}
