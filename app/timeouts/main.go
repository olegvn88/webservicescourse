package main

import (
	"fmt"
	"time"
)

func setTimeouts() {
	// при 1 выполнится таймаут, при 3 - выполнится операция
	timer := time.NewTimer(2 * time.Second)
	select {
	case <-timer.C:
		fmt.Println("timer.C timeout happened")
	case <-time.After(time.Minute):
		// пока не выстрелит - не соберётся сборщиком мусора
		fmt.Println("time.After timeout happened")
	case result := <-longSQLQuery():
		// освобождет ресурс
		if !timer.Stop() {
			<-timer.C
		}
		fmt.Println("operation result:", result)

	}
}

func longSQLQuery() chan bool {
	ch := make(chan bool, 1)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- true
	}()
	return ch
}

func main() {
	setTimeouts()
}
