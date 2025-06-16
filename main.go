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

var templates = template.Must(template.ParseFiles("templates/index.html", "templates/confirmation.html", "templates/error.html"))

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
	data := struct {
		AppName          string
		RemainingTickets uint
		TotalTickets     uint
	}{
		AppName:          appName,
		RemainingTickets: remainingTickets,
		TotalTickets:     totalTickets,
	}
	mutex.Unlock()

	if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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

	// Render confirmation page
	data := struct {
		FirstName       string
		LastName        string
		Email           string
		NumberOfTickets uint
		Remaining       uint
	}{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		NumberOfTickets: userTickets,
		Remaining:       remainingTickets,
	}

	if err := templates.ExecuteTemplate(w, "confirmation.html", data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func renderError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	err := templates.ExecuteTemplate(w, "error.html", struct{ Message string }{Message: message})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
