package main

import (
	"fmt"
)

func main() {
	appName := "ReserGo"
	const noOfTickets int = 100
	var remainingTickets uint = 100

	fmt.Printf("Welcome to %v !\n", appName)
	fmt.Printf("We have total %v tickets and %v tickets are available\n", noOfTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend the conference")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve your booking confirmation on %v email", firstName, lastName, userTickets, email)

}
