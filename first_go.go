package main

import "fmt"

func main() {
	account := Account{
		Holder:  "Roger",
		Number:  1001,
		Balance: 15000,
	}

	fmt.Printf("Account holder: %s\n", account.Holder)
	fmt.Printf("Account number: %d\n", account.Number)
	fmt.Printf("Original Balance: %d cents\n", account.Balance)

	newBalance, err := withdrawal(account.Balance, 5000)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		account.Balance = newBalance
		fmt.Printf("Withdrawal successful!\nUpdated balance: %d cents\n", account.Balance)
	}
}

type Account struct {
	Holder  string
	Number  int
	Balance int
}

func deposit(balance, amount int) int {
	// since both params are same type just specify once
	// return type doesn't need parensthesis on single type
	return balance + amount
}

func withdrawal(balance, amount int) (int, error) {
	if amount <= 0 {
		return balance, fmt.Errorf("invalid withdrawal amount")
	}

	if amount > balance {
		return balance, fmt.Errorf("insufficient funds")
	}

	return balance - amount, nil
}
