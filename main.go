package main

import (
	"ReserGo/helper"
	"fmt"
)

var appName = "ReserGo"

const noOfTickets int = 100

var remainingTickets uint = 100
var bookings = make([]UserData, 0)

type UserData struct {
	firstName              string
	lastName               string
	email                  string
	numberOftickets        uint
	isOptedInForNewsLetter bool
}

func main() {

	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			printFirstNames()

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
			if userTickets > remainingTickets {
				fmt.Printf("Sorry! We have only %v tickets available, so you can't book %v tickets\n", remainingTickets, userTickets)
			}
		}

	}
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", appName)
	fmt.Printf("We have total %v tickets and %v tickets are available\n", noOfTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend the conference")
}

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("All the bookings: \n%v\n", firstNames)
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:              firstName,
		lastName:               lastName,
		email:                  email,
		numberOftickets:        userTickets,
		isOptedInForNewsLetter: false,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve your booking confirmation on %v email\n", firstName, lastName, userTickets, email)
	fmt.Printf("Now we have %v tickets available\n", remainingTickets)
}
