package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTicket int = 50
	var remianingTicket uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v are still availabe.\n", conferenceTicket, remianingTicket)
	fmt.Println("Get you Tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for thier details to enter.
		fmt.Println("Enter your first Name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last Name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address")
		fmt.Scan(&email)

		fmt.Printf("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		remianingTicket = remianingTicket - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf(" Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remianingTicket, conferenceName)

		fmt.Printf("These are all our bookings: %v\n", bookings)

	}

}
