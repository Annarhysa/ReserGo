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

	greetUser(appName, noOfTickets, remainingTickets)

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// isValidCity := city != "Singapore" || city != "London"

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(remainingTickets, userTickets, firstName, lastName, email, bookings)

			printFirstNames(bookings)

			if remainingTickets == 0 {
				fmt.Println("All tickets are sold out!")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered is not valid")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
			fmt.Printf("Sorry! We have only %v tickets available, so you can't book %v tickets\n", remainingTickets, userTickets)
		}

	}
}

func greetUser(confName string, noOfTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total %v tickets and %v tickets are available\n", noOfTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend the conference")
}

func printFirstNames(bookings []string) {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("All the bookings: \n%v\n", firstNames)
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, firstName string, lastName string, email string, bookings []string) {
	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve your booking confirmation on %v email\n", firstName, lastName, userTickets, email)
	fmt.Printf("Now we have %v tickets available\n", remainingTickets)
}
