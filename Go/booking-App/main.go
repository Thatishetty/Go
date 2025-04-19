package main

import (
	"fmt"
	"strings"
)

const conferanceTickets int = 50

var conferanceName = "Go Conferance"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInputs()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstName()
			fmt.Printf("These are First names of Bookings   : %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("The Tickets for %v are Booked Out come Back Next year.\n", conferanceName)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your first name or last name is too Short")
			}
			if !isValidEmail {
				fmt.Println("Enter Corrrect Email Format with @")
			}
			if !isValidTicketNumber {
				fmt.Println("no of tickets you entered is Invalid")
			}
		}
	}

}

func greetUsers() {
	fmt.Printf("*-------------------------------------------------------------*\n")
	fmt.Printf("Welcome to %v Booking Application\n", conferanceName)
	fmt.Printf("We have total %v tickets and %v are still available\n", conferanceTickets, remainingTickets)
	fmt.Println("Get your tickets here")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter User First Name : ")
	fmt.Scan(&firstName)

	fmt.Println("Enter User Last Name : ")
	fmt.Scan(&lastName)

	fmt.Println("Enter User Email Address :")
	fmt.Scan(&email)

	fmt.Println("Enter No of Tickets :")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - remainingTickets


	var userData 

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("-------------------------------------------------------------\n")
	fmt.Printf("Thank You %v %v for Booking %v tickets . You will receive a conformation mail to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Tickets left for %v\n", remainingTickets, conferanceName)
}

// userName = "Tom"
// userTickets = 7
// city := London
// call firstnames func
// switch city{
// case "London" :
// 	//Execute code for london
// case "New York" :
// 	// some code
// default:
// 	fmt.Printlv("The city you entered is Invalid")

// fmt.Printf("Booked by %v\n", bookings)
// fmt.Printf("The First Booking  %v\n", bookings[0])
// fmt.Printf("Slice Length  %v\n", len(bookings))
// fmt.Printf("Slice type %T\n", bookings)
// // fmt.Printf("**-------------------------------------------------------------**\n")

// fmt.Printf("conferanceTickets Type is %T,conferanceName Type is %T.\n", conferanceTickets, conferanceName)
