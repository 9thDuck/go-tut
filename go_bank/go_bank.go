package main

import "fmt"

func main() {
	accountBalance := 1000.0
	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		var choice int
		fmt.Scan(&choice)

		// if choice == 1 {
		// 	fmt.Printf("Your account balance is %.2f\n\n", accountBalance)
		// } else if choice == 2 {
		// 	deposiAmount := getFloat64UserInput("Enter deposit amount:")
		// 	if deposiAmount <= 0 {
		// 		fmt.Printf("Invalid deposit amount \"%.2f\". Deposit amount must be greater than 0.", deposiAmount)
		// 		continue
		// 	}

		// 	accountBalance += deposiAmount
		// 	fmt.Printf("Updated account balance is %.2f\n", accountBalance)
		// } else if choice == 3 {
		// 	withdrawlAmount := getFloat64UserInput("Enter withdrawl amount:")
		// 	if withdrawlAmount > accountBalance {
		// 		fmt.Printf("Withdrawl amount \"%.2f\" is greater than account balance \"%.2f\". Withdrawl amount must be less than account balance\n\n", withdrawlAmount, accountBalance)
		// 		continue
		// 	} else if withdrawlAmount <= 0 {
		// 		fmt.Printf("Withdrawl amount \"%.2f\" is invalid input. Withdrawl amount must be greater than 0.\n\n", withdrawlAmount)
		// 		continue
		// 	}

		// 	accountBalance -= withdrawlAmount
		// 	fmt.Printf("Updated account balance is %.2f\n\n", accountBalance)
		// } else if choice == 4 {
		// 	fmt.Println("Thanks for choosing Go bank!")
		// 	break
		// }

		switch choice {
		case 1:
			fmt.Printf("Your account balance is %.2f\n\n", accountBalance)
		case 2:
			deposiAmount := getFloat64UserInput("Enter deposit amount:")
			if deposiAmount <= 0 {
				fmt.Printf("Invalid deposit amount \"%.2f\". Deposit amount must be greater than 0.", deposiAmount)
				continue
			}
			accountBalance += deposiAmount
			fmt.Printf("Updated account balance is %.2f\n\n", accountBalance)
		case 3:
			withdrawlAmount := getFloat64UserInput("Enter withdrawl amount:")
			if withdrawlAmount > accountBalance {
				fmt.Printf("Withdrawl amount \"%.2f\" is greater than account balance \"%.2f\". Withdrawl amount must be less than or equal to account balance\n\n", withdrawlAmount, accountBalance)
				continue
			} else if withdrawlAmount <= 0 {
				fmt.Printf("Withdrawl amount \"%.2f\" is invalid input. Withdrawl amount must be greater than 0.\n\n", withdrawlAmount)
				continue
			}
			accountBalance -= withdrawlAmount
			fmt.Printf("Updated account balance is %.2f\n\n", accountBalance)
		case 4:
			fmt.Println("Thanks for choosing Go bank!")
			return

	}
}

func getFloat64UserInput(label string) (input float64) {
	fmt.Print(label)
	fmt.Scan(&input)
	return input
}
