package main

import (
	"fmt"
	"runtime"
	"time"
)

const goroutinesNum = 3

func startWorker(workerNum int, in <-chan string) {
	for input := range in {
		fmt.Printf(formatWork(workerNum, input))
		runtime.Gosched()
	}
	printFinishWork(workerNum)
}

func printFinishWork(i int) {
	fmt.Println(i)
}

func formatWork(i int, s string) string {

	return s
}

func main() {
	workerInput := make(chan string, 2)
	for i := 0; i < goroutinesNum; i++ {
		go startWorker(i, workerInput)
	}

	months := []string{"Январь", "Февраль", "Март",
		"Апрель", "Май", "Июнь",
		"Июль", "Август", "Сентябрь",
		"Октябрь", "Ноябрь", "Декабрь",
	}

	for _, monthName := range months {
		workerInput <- monthName
	}
	close(workerInput)

	time.Sleep(time.Millisecond)
}
