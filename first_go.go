package main

import "fmt"

func main() {
    accountHolder := "Roger"
    accountNum := 1001
    balance := 15000

    fmt.Printf("Account holder: %s\n", accountHolder)
    fmt.Printf("Account number: %d\n", accountNum)
    fmt.Printf("Original Balance: %d cents\n", balance)
    
    newBalance, err := withdrawal(balance, 5000)
    
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Withdrawal successful! New balance: %d cents\n", newBalance)
    }
    

    newBalance, err = withdrawal(balance, 20000)

    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Withdrawal successful! New balance: %d cents\n", newBalance)    
    }
}


func deposit(balance, amount int) int {
    // since both params are same type just specify once
    // return type doesn't need parensthesis on single type
    return balance + amount
}

func withdrawal(balance, amount int) (int, error) {
    if amount <= 0 {
        return  balance, fmt.Errorf("invalid withdrawal amount")
    }

    if amount > balance {
        return  balance, fmt.Errorf("insufficient funds")
    }

    return balance - amount, nil
}