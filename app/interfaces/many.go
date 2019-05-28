package main

import (
	"fmt"
)

//у нас есть кошелек, он реализкует метод Pay, у нас есть некоторое количество денег в нем

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Not enough money")
	}
	w.Cash -= amount
	return nil
}

type Card struct {
	Balance    int
	ValidUntil string
	CardHolder string
	CVV        string
	Number     string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("Not enough money")
	}
	c.Balance -= amount
	return nil
}

type ApplePay struct {
	Money   int
	AppleID string
}

func (a *ApplePay) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("Not enough money")
	}

	a.Money -= amount
	return nil
}

type Payer interface {
	Pay(int) error
}

//func Buy(p Payer) {
//	err := p.Pay(10)
//	if err != nil {
//		fmt.Printf("Error during payment %T: %v\n\n", p, err)
//		return
//	}
//	fmt.Printf("Thank you for the purchase %T\n\n", p)
//}

func Buy(p Payer) {
	switch p.(type) {
	case *Wallet:
		fmt.Println("Cash payment?")
	case *Card:
		plasticCard, ok := p.(*Card)
		if !ok {
			fmt.Println("failed to convert to type Card")
		}
		fmt.Println("Insert the card", plasticCard.CardHolder)
	default:
		fmt.Println("Something is new!")
	}

	err := p.Pay(10)
	if err != nil {
		panic(err)
	}
}

func main() {
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)

	var myMoney Payer
	myMoney = &Card{Balance: 100, CardHolder: "onesterov"}
	Buy(myMoney)

	myMoney = &ApplePay{Money: 9}
	Buy(myMoney)
}
