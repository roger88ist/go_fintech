package main

import "fmt"

func main() {
    accountHolder := "Roger"
    accountNum := 1001
    balance := 15000
    depositAmount := 5000
    newBalance := deposit(balance, depositAmount)

    fmt.Printf("Account holder: %s\n", accountHolder)
    fmt.Printf("Account number: %d\n", accountNum)
    fmt.Printf("Original Balance: %d cents\n", balance)
    fmt.Printf("Updated Balance: %d cents\n", newBalance)
}


func deposit(balance, amount int) int {
    // since both params are same type just specify once
    //  return type doesn't need parensthesis on single type
    return balance + amount
}