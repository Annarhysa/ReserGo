package main

import (
	"ReserGo/helper"
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	appName       = "ReserGo"
	totalTickets  = 100
	minNameLength = 2
)

var (
	remainingTickets uint = totalTickets
	bookings              = make([]UserData, 0)
	mutex            sync.Mutex
	wg               sync.WaitGroup
)

type UserData struct {
	firstName              string
	lastName               string
	email                  string
	numberOfTickets        uint
	isOptedInForNewsLetter bool
}

func main() {
	greetUser()

	for remainingTickets > 0 {
		firstName, lastName, email, userTickets, optedIn := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if !isValidName {
			fmt.Println("Error: First name and last name must be at least 2 characters long.")
			continue
		}
		if !isValidEmail {
			fmt.Println("Error: Email address is invalid. Please enter a valid email.")
			continue
		}
		if !isValidTicketNumber {
			fmt.Printf("Error: Number of tickets must be between 1 and %d.\n", remainingTickets)
			continue
		}

		bookTicket(userTickets, firstName, lastName, email, optedIn)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		printFirstNames()

		if remainingTickets == 0 {
			fmt.Println("All tickets are sold out! Thank you for your bookings.")
			break
		}

		fmt.Println("Do you want to book more tickets? (yes/no)")
		if !askForYes() {
			break
		}
	}

	wg.Wait()
	fmt.Println("Exiting application. Have a nice day!")
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", appName)
	fmt.Printf("We have a total of %v tickets and %v tickets are available\n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend the conference")
}

// printFirstNames prints the first names of all users who have booked tickets.
func printFirstNames() {
	mutex.Lock()
	defer mutex.Unlock()

	firstNames := make([]string, len(bookings))
	for i, booking := range bookings {
		firstNames[i] = booking.firstName
	}
	fmt.Printf("Current bookings by: %v\n", strings.Join(firstNames, ", "))
}

// getUserInput collects user input and validates basic formatting.
func getUserInput() (string, string, string, uint, bool) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your first name: ")
	firstName := readNonEmptyString(reader)

	fmt.Print("Enter your last name: ")
	lastName := readNonEmptyString(reader)

	fmt.Print("Enter your email: ")
	email := readNonEmptyString(reader)

	var userTickets uint
	for {
		fmt.Print("Enter number of tickets: ")
		_, err := fmt.Scan(&userTickets)
		if err != nil || userTickets == 0 {
			fmt.Println("Please enter a valid positive number for tickets.")
			// clear input buffer
			reader.ReadString('\n')
			continue
		}
		break
	}

	fmt.Print("Do you want to opt-in for our newsletter? (yes/no): ")
	optedIn := askForYes()

	return firstName, lastName, email, userTickets, optedIn
}

// readNonEmptyString reads a non-empty trimmed string from the reader.
func readNonEmptyString(reader *bufio.Reader) string {
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if len(input) > 0 {
			return input
		}
		fmt.Print("Input cannot be empty. Please enter again: ")
	}
}

// askForYes returns true if user inputs 'yes' (case insensitive), false otherwise.
func askForYes() bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "yes" || input == "y" {
			return true
		} else if input == "no" || input == "n" {
			return false
		} else {
			fmt.Print("Please enter 'yes' or 'no': ")
		}
	}
}

// bookTicket safely updates remaining tickets and bookings.
func bookTicket(userTickets uint, firstName, lastName, email string, optedIn bool) {
	mutex.Lock()
	defer mutex.Unlock()

	remainingTickets -= userTickets

	userData := UserData{
		firstName:              firstName,
		lastName:               lastName,
		email:                  email,
		numberOfTickets:        userTickets,
		isOptedInForNewsLetter: optedIn,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive your booking confirmation at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Tickets remaining: %v\n", remainingTickets)
}

// sendTicket simulates sending ticket asynchronously.
func sendTicket(userTickets uint, firstName, lastName, email string) {
	defer wg.Done()

	time.Sleep(5 * time.Second) // reduced sleep for demo purposes
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	fmt.Println("#########################")
	fmt.Printf("Sending ticket:\n%v \nto email address: %v\n", ticket, email)
	fmt.Println("#########################")
}
