package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
	
		// get user details
		firstName, lastName, email, userTickets := getuserDetails()

		// User input validation
		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTickets {
			
			bookTickets(&remainingTickets, userTickets, &bookings, firstName, lastName, email, conferenceName)
	
			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are : %v \n", firstNames)
	
			if remainingTickets == 0 {
				fmt.Println("The Conference is sold out. Come back next year")
				//end program
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address doesnt contain @ sign")
			}
			if !isValidTickets {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	
		
	}
	
}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %s booking application. \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %d are still available. \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		fullName := strings.Fields(booking)
		firstNames = append(firstNames, fullName[0])
	}
	//fmt.Printf("The first names of bookings are : %v \n", firstNames)
	return firstNames
}



func getuserDetails() (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(remainingTickets *uint, userTickets uint, bookings *[]string, firstName string, lastName string, email string, conferenceName string) {
	*remainingTickets = *remainingTickets - userTickets
	*bookings = append(*bookings, firstName + " " + lastName)

	fmt.Printf("Thank you %s for booking %d tickets. You will recieve a confirmation email at %v \n", firstName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", *remainingTickets, conferenceName)
}