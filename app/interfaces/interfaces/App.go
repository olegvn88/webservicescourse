package main

import "fmt"

type Payer interface {
	Pay(int) error
}

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("It is not enough money")
	}
	w.Cash -= amount
	return nil
}

func Buy(p Payer) {
	err := p.Pay(10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Thank you for the your purchase with %T\n\n", p) // %T тип переданного аргумента
}

func main() {
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)
}
