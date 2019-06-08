package main

import (
	"fmt"
)

func main() {
	chi := make(chan int, 2)
	go func(in chan int) {
		val := <-in
		fmt.Println("GO: get from chan", val)
		fmt.Println("GO: after read from chan")
	}(chi)
	chi <- 42
	chi <- 43
	fmt.Println("MAIN: after put to the chan")
	//fmt.Scanln()
	//fmt.Println(<-getData("Oleg"))
	for i := range getData("") {
		fmt.Println(" ", i)
	}
	//fmt.Println(<-getData(""))

}

func getData(val string) chan int {
	chanData := make(chan int, 0)

	go func(out chan<- int) {
		for i := 0; i <= 10; i++ {
			out <- i
		}
		close(out) // close channel / if
	}(chanData)

	//for i := range chanData {
	//	fmt.Println(i)
	//}
	return chanData
}

//func getData(val string) chan string {
//	chanData := make(chan string, 0)
//
//	go func(out chan<- string) {
//		for i := 0; i < 10; i++ {
//			out <- strconv.Itoa(i)
//		}
//		close(out)
//	}(chanData)
//	return chanData
//}
