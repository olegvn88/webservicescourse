package main

import (
	"fmt"
	"time"
)

func tickerExample() {
	ticker := time.NewTicker(time.Second)
	i := 0
	for tickTime := range ticker.C {
		i++
		fmt.Println("step", i, "time", tickTime)
		if i >= 5 {
			ticker.Stop()
			break
		}
	}
}

func ticker2() {

	c := time.Tick(time.Second)
	i := 0
	for tickTime := range c {
		i++
		fmt.Println("step", i, "time", tickTime)
		if i >= 5 {
			break
		}
	}
}

func main() {
	tickerExample()
	ticker2()
}
