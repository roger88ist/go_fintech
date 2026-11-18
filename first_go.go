package main

import (
	"errors"
	"fmt"
)

func main() {
	account := Account{
		Holder:  "Roger",
		Number:  1001,
		Balance: 15000,
	}

    displayAccount(account)

	err := account.Withdraw(5000)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Withdrawal successful! Balance: %d cents\n", account.Balance)
	}

	err = account.Withdraw(20000)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Withdrawal successful! Balance: %d cents\n", account.Balance)
	}

	err = account.Withdraw(-500)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Withdrawal successful! Balance: %d cents\n", account.Balance)
	}

	fmt.Printf("Final balance: %d cents\n", account.Balance)
}

type Account struct {
	Holder  string
	Number  int
	Balance int
}

func (a *Account) Withdraw(amount int) error {
	if amount <= 0 {
		return errors.New("invalid withdrawal amount")
	}

	if amount > a.Balance {
		return errors.New("insufficient funds")
	}

	a.Balance -= amount
	return nil
}

func (a Account) PrintInfo() {
    fmt.Printf("Account holder: %s\n", a.Holder)
    fmt.Printf("Account number: %d\n", a.Number)
    fmt.Printf("Account Balance: %d cents\n", a.Balance)
}

type AccountPrinter interface {
    PrintInfo()
}

func displayAccount(a AccountPrinter) {
    a.PrintInfo()
}