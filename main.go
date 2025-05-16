// The first statement of go must be package
package main

import "fmt" //for print function

func main() {
	var conferenceName = "Go to Conference"
	const conferenceTickets = 50
	var remainingTickets = 46
	//type print
	fmt.Printf("conferenct Tickets is %T\n", conferenceTickets)

	fmt.Println("Welcome to our", conferenceName, " booking application")
	fmt.Println("we have total tickets", conferenceTickets, "Remaining ticket", remainingTickets)
	fmt.Println("Get your tickets to attend")
	//data types
	var userName string
	var userTickets int

	userName = "tom"
	userTickets = 14
	fmt.Printf("user %v booked %v tickets .\n", userName, userTickets)

}
