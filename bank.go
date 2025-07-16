package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

//the above creates a global value

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}
	// the above if statment is how to handle potential errors in golang
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}

	return balance, nil
}

//the underscore tell go to not worry about the error if one happens

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

// file 0644 is the linux way of giving permissions to a file
func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------")
		return
	}
	//return can end the program but panic("can't continue makes the program look more like an error and crash")

	fmt.Println("Welcome to Go Bank!")
	//for is the only type of loop in golang
	// the below is a loop that starts at zero
	// then states that i can't be greater than 2
	//and finally the last part increases the runs of the loop
	//below currently the for loop will go on forever unless closed
	// here below is how to add a set number of loops
	//i := 0; i < 2; i++

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		fmt.Println("Your choice:", choice)

		//if you have an application where you need
		// to break out of a loop and continue the application then you must use
		// if else statements will be the only way to go.
		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				//by adding the return the code will stop after this point.
				//return
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				//by adding the return the code will stop after this point.
				//return
				//skips the rest of the loop
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. Your withdraw can't exceed your balance.")
				continue
			}

			accountBalance += withdrawAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Have a good day")
			fmt.Println("Thanks for choosing my bank")
			return
		}

	}
}
