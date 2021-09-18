package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	done := make(chan struct{})
	go func() {
		work()
		done <- struct{}{}
	}()
	// join point needed
	<-done
	fmt.Println("elapse time:", time.Since(now))
	fmt.Println("done waiting, main exits!")
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("printing some stuff")
}
