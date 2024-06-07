package main

import "fmt"

func main() {
	conferenceName := "DolBy"

	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("wellcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v and %v still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string

	var lastName string

	var email string

	var userTickets int

	fmt.Println("enter the user first name:")
	fmt.Scan(&firstName)
	fmt.Println("enter the user last name:")
	fmt.Scan(&lastName)
	fmt.Println("enter the user email id :")
	fmt.Scan(&email)
	fmt.Println("enter the no of tickets :")
	fmt.Scan(&userTickets)
	fmt.Printf("Thank you %v %v for booking of %v tickets. you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("%v  remaining tickets for  %v\n", remainingTickets, conferenceName)

}
