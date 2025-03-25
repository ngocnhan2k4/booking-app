package main

import "fmt"


func main(){
	var conferenceName string = "GopherCon"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Println("Welcome to", conferenceName)
	fmt.Println("Number of tickets left:", conferenceTickets)
	fmt.Println("Number of tickets remaining:", remainingTickets)
}