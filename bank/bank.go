package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	println("Welcome to Gobber Bank")
	println("What do you want to do?")
	println("1. Check balance")
	println("2. Deposit money")
	println("3. Withdraw money")
	println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is:", accountBalance)
	} else if choice == 2 {
		fmt.Print("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0.")
			return
		}

		accountBalance += depositAmount
		fmt.Println("Balance updated! New amount: ", accountBalance)
	} else if choice == 3 {
		fmt.Print("Withdrawal amount: ")
		var withdrawalAmount float64
		fmt.Scan(&withdrawalAmount)

		if withdrawalAmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0.")
			return
		} else if withdrawalAmount > accountBalance {
			fmt.Println("You dont have enough money!")
			return
		}

		accountBalance -= withdrawalAmount
		fmt.Println("Balance updated! New amount: ", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}

	fmt.Println("Your choice", choice)

}
