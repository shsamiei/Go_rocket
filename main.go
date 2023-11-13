package main

import "fmt"

// where to start ? main function 
func main(){
	var ConferenceName string = "Go Conference"
	const ConferenceTickets int = 50
	var remainingTickets int = 50


	fmt.Printf("We can use %T for show the type of var, for example in here I want to show the type of Conference Ticket", ConferenceTickets)


	// fmt.Println("Welcome to the Nexaloak", ConferenceName ,"Society")
	fmt.Printf("Welcome to the Nexaloak %v Society", ConferenceName)
	fmt.Println("Here you can book one node in Nexaloak Blockchain")
	// fmt.Println("There are ", ConferenceTickets, "Tickets and ", remainingTickets, "are still available")
	fmt.Printf("There are %v Tickets and %v  are still available", ConferenceTickets, remainingTickets)


	var userName string
	var userTicket int

	fmt.Println(userTicket)
	// ask user for name 

	userName = "Tom"
	fmt.Println(userName)




}