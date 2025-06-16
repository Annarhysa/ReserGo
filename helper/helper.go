package helper

import (
	"regexp"
	"strings"
)

// ValidateUserInput validates user input for name, email, and ticket number.
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(strings.TrimSpace(firstName)) >= 2 && len(strings.TrimSpace(lastName)) >= 2
	isValidEmail := isEmailValid(email)
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

// isEmailValid uses a regex to validate email format.
func isEmailValid(email string) bool {
	email = strings.TrimSpace(email)
	if len(email) < 6 {
		return false
	}
	// Simple regex for email validation
	regex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
