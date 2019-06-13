package main

import "fmt"

func main() {
	//fmt.Println("Hello man")

	go printHello()
	fmt.Scanln()
}

func printHello() chan int {
	ch := make(chan int, 4)

	for i := 0; i < 4; i++ {
		ch <- i
		fmt.Print(<-ch)
	}

	return ch
}
