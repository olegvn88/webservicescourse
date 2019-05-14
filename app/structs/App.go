package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go Count("Hello", c)

	for {
		msg := <-c
		fmt.Println(msg)
	}
}

func Count(s string, c chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, s)
		c <- s
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
