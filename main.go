package main

import (
	"fmt"
	"strings"
)

func main() {
	appName := "ReserGo"
	const noOfTickets int = 100
	var remainingTickets uint = 100
	bookings := []string{}

	fmt.Printf("Welcome to %v !\n", appName)
	fmt.Printf("We have total %v tickets and %v tickets are available\n", noOfTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend the conference")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("Sorry! We have only %v tickets available, so you can't book %v tickets\n", remainingTickets, userTickets)
			break
		}

		remainingTickets = remainingTickets - userTickets

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve your booking confirmation on %v email\n", firstName, lastName, userTickets, email)
		fmt.Printf("Now we have %v tickets available\n", remainingTickets)

		firstNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("All the bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("All tickets are sold out!")
			break
		}
	}
}
