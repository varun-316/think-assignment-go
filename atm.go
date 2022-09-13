package main

import "fmt"

func ATMMenu() {
	a := 0
	for a != -1 {
		fmt.Print("1. Details \n2. Deposit\n3. Withdraw \n")
		fmt.Scanln(&a)
		switch a {
			case 1: 
				fmt.Println("Details")
			case 2: 
				fmt.Println("Deposit")
			case 3:
				fmt.Println("Withdraw")
			default:
				fmt.Println("Enter a valid option")
		}
	}
}

func GetDetails() {

}

func DepositMoney() {

}

func WithdrawMoney() {
	
}