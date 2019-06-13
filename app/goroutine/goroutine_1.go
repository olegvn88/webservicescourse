package main

import (
	"fmt"
	"time"
)

func main() {
	go hello()
	time.Sleep(1 * time.Second)
	for i := 0; i < 1; i++ {
		fmt.Println("main function")
	}
	//fmt.Scanln()
}

func hello() {
	//for i := 0; i < 10; i++ {
	fmt.Println("Hello world goroutine")
	//}
}
