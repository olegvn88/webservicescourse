package main

import (
	"fmt"
)

func playWithChannels() chan int {
	ch1 := make(chan int)
	go func(in chan int) {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(in)
	}(ch1)

	return ch1
}

func multiplexing() {
	channel1 := make(chan int, 1)
	channel2 := make(chan int)

	channel1 <- 1
	//channel2 <- 2

	select {
	case val := <-channel1:
		fmt.Println(val)
	case channel2 <- 2:
		fmt.Println("put value to chan 2")
	default:
		fmt.Println("default case")
	}
}

func multiplexingINLOOP() {
	channel1 := make(chan int, 2)
	channel2 := make(chan int, 2)

	channel1 <- 1
	channel1 <- 2

	channel2 <- 3
	//channel2 <- 2
LOOP:
	for {
		select {
		case val := <-channel1:
			fmt.Println("chan1 val", val)
		case val2 := <-channel2:
			fmt.Println("chan2 val", val2)
		default:
			break LOOP
		}
	}
}

//завершение внешней функции, используя канал отмены

func cancelChfunc() {
	cancelCh := make(chan struct{})
	dataCh := make(chan int)

	go func(cancelCh chan struct{}, dataCh chan int) {
		val := 0
		for {
			select {
			case <-cancelCh:
				return
			case dataCh <- val:
				val++
			}
		}

	}(cancelCh, dataCh)

	for curVal := range dataCh {
		fmt.Println("read", curVal)
		if curVal > 3 {
			fmt.Println("send cancel")
			cancelCh <- struct{}{}
			break
		}
	}
}

func main() {
	//multiplexing()
	//multiplexingINLOOP()
	cancelChfunc()
}
