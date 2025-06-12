package main

import "fmt"

func main() {
	var accountBalance = 1000.0
	fmt.Println("Welcome to Go Bank!")
	//for is the only type of loop in golang
	// the below is a loop that starts at zero
	// then states that i can't be greater than 2
	//and finally the last part increases the runs of the loop
	//below currently the for loop will go on forever unless closed
	// here below is how to add a set number of loops
	//i := 0; i < 2; i++

	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		fmt.Println("Your choice:", choice)

		if choice == 1 {
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 {
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
		} else if choice == 3 {
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
				return
			}

			accountBalance += withdrawAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
		} else {
			fmt.Println("Have a good day")
			//return exits if you are in a for loop or not but it closes the whole program
			//return
			//break closes just the for loop but everything else runs
			break
		}
	}

	fmt.Println("Thanks for choosing my bank")
}
