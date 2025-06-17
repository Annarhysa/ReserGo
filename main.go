package main

import (
	"ReserGo/helper"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const (
	appName      = "ReserGo"
	totalTickets = 100
)

var (
	remainingTickets uint = totalTickets
	bookings              = make([]UserData, 0)
	mutex            sync.Mutex
)

type UserData struct {
	FirstName              string
	LastName               string
	Email                  string
	NumberOfTickets        uint
	IsOptedInForNewsLetter bool
}

// Unified struct for passing data to templates
type TemplateData struct {
	AppName          string
	RemainingTickets uint
	TotalTickets     uint
	Message          string

	// Fields for confirmation page
	FirstName       string
	LastName        string
	Email           string
	NumberOfTickets uint
	Remaining       uint
}

var templates = template.Must(template.ParseFiles(
	"templates/layout.html",
	"templates/index.html",
	"templates/confirmation.html",
	"templates/error.html"))

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/book", bookHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Printf("Starting %s server on http://localhost:8080\n", appName)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// homeHandler serves the booking form
func homeHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	data := TemplateData{
		AppName:          appName,
		RemainingTickets: remainingTickets,
		TotalTickets:     totalTickets,
	}
	mutex.Unlock()

	if err := templates.ExecuteTemplate(w, "layout.html", data); err != nil {
		log.Printf("Template execution error in homeHandler: %v", err)
		renderError(w, "Internal Server Error")
		return
	}
}

// bookHandler processes the booking form submission
func bookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()
	if err != nil {
		renderError(w, "Invalid form data")
		return
	}

	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	email := r.FormValue("email")
	ticketsStr := r.FormValue("tickets")
	newsletter := r.FormValue("newsletter") == "on"

	userTickets64, err := strconv.ParseUint(ticketsStr, 10, 32)
	if err != nil || userTickets64 == 0 {
		renderError(w, "Please enter a valid number of tickets")
		return
	}
	userTickets := uint(userTickets64)

	mutex.Lock()
	defer mutex.Unlock()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if !isValidName {
		renderError(w, "First name and last name must be at least 2 characters long")
		return
	}
	if !isValidEmail {
		renderError(w, "Email address is invalid")
		return
	}
	if !isValidTicketNumber {
		renderError(w, fmt.Sprintf("Number of tickets must be between 1 and %d", remainingTickets))
		return
	}

	// Book tickets
	remainingTickets -= userTickets
	userData := UserData{
		FirstName:              firstName,
		LastName:               lastName,
		Email:                  email,
		NumberOfTickets:        userTickets,
		IsOptedInForNewsLetter: newsletter,
	}
	bookings = append(bookings, userData)

	// Render confirmation page with expected data fields
	data := TemplateData{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		NumberOfTickets: userTickets,
		Remaining:       remainingTickets,
		AppName:         appName,
	}

	if err := templates.ExecuteTemplate(w, "layout.html", data); err != nil {
		log.Printf("Template execution error in bookHandler: %v", err)
		renderError(w, "Internal Server Error")
		return
	}
}

// renderError renders the error page with the given message
func renderError(w http.ResponseWriter, message string) {
	log.Printf("renderError called with message: %s", message)
	data := TemplateData{
		AppName: appName,
		Message: message,
	}

	w.WriteHeader(http.StatusBadRequest)
	err := templates.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		log.Printf("Error rendering error template: %v", err)
		// Do NOT call http.Error again because headers are already sent
	}
}
