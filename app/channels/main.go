package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	done := make(chan bool)
	wg.Add(1)
	go hello(wg, done)
	<-done
	fmt.Println("main functions")
	wg.Wait()
}

func hello(wg *sync.WaitGroup, done chan bool) {
	defer wg.Done()
	fmt.Println("Hello world goroutine")
	done <- true
}
