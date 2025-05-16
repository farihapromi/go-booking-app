package main

import "fmt"

func main() {
	var conferenceName = "Go to Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 46

	// Print types
	fmt.Printf("conferenceTickets is of type: %T\n", conferenceTickets)

	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Println("We have total tickets:", conferenceTickets, "Remaining tickets:", remainingTickets)
	fmt.Println("Get your tickets to attend")

	// Input data
	var userName string
	var lastName string
	var userTickets uint
	var email string

	// Read user input
	fmt.Println("Enter first name:")
	fmt.Scan(&userName)

	fmt.Println("Enter last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	// Update remaining tickets
	remainingTickets = remainingTickets - userTickets

	// Confirmation message
	fmt.Printf("Thank you %v %v for booking %v tickets.\n", userName, lastName, userTickets)
	fmt.Printf("You will receive confirmation at %v\n", email)
	fmt.Printf("Tickets remaining: %v\n", remainingTickets)

	// Append to bookings slice
	var bookings []string
	bookings = append(bookings, userName+" "+lastName)

	fmt.Printf("Bookings: %v\n", bookings)
	fmt.Printf("Number of bookings: %v\n", len(bookings))
}
