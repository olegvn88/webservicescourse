package main

import (
	"fmt"
	"strconv"
)

type Wallet struct {
	Cash int
}

type Payer interface {
	Pay(int) error
}

func (w *Wallet) String() string {
	return "Wallet, which has " + strconv.Itoa(w.Cash) + " money"
}

func main() {
	myWallet := &Wallet{Cash: 100}
	fmt.Printf("Raw payment : %#v\n", myWallet)   // полное представление структуры
	fmt.Printf("Payment method : %s\n", myWallet) // строка
}
