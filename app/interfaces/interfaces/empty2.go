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

func Buy(in interface{}) {
	var p Payer
	var ok bool
	if p, ok = in.(Payer); !ok {
		fmt.Printf("%T is not means of payment\n\n", in)
		return
	}
	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Error during of payment %T: %v\n\n", p, err)
	}
	fmt.Printf("Thank you for the purchase via %T\n\n", p)
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Not enough money")
	}
	w.Cash -= amount
	return nil
}

func main() {
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)
	Buy([]int{1, 2, 3})
	Buy(3.14)
}
