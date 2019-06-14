package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(4)
	go printHello(wg)
	go prinData(4, wg)
	go prinData(2, wg)
	go printRuneLetters(4, wg)
	wg.Wait()

}

func prinData(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch := make(chan rune, 1)
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
		ch <- b[i]
		fmt.Printf("%c", <-ch)
	}
	fmt.Println(rand.Intn(len(letterRunes)))
}

func printHello(wg *sync.WaitGroup) chan int {
	defer wg.Done()
	ch := make(chan int, 1)

	for i := 0; i < 4; i++ {
		ch <- i
		func() {

			fmt.Print(" ", <-ch, " ")
		}()
	}

	return ch
}

func printRuneLetters(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	data := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	arr := make([]rune, n)
	ch1 := make(chan rune, 1)
	for i := range arr {
		arr[i] = data[rand.Intn(len(data))]
		ch1 <- arr[i]
		fmt.Printf("%c", <-ch1)
	}
}
