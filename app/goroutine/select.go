package main

import (
	"fmt"
)

func selectFunc() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)

	ch1 <- 1
	ch1 <- 2
	ch2 <- 3
LOOP:
	for {

		select {
		case v1 := <-ch1:
			fmt.Println("chan val1", v1)
		case v2 := <-ch2:
			fmt.Println("chan val2", v2)
		default:
			break LOOP // LOOP - метка цикла
		}
	}
}

func main() {
	selectFunc()
}
