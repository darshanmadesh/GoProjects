package main

import (
	"fmt"
	"strings"
)

type userData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	var bookings = make([]userData, 0)

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		// get user details
		var userInfo = getuserDetails()

		// User input validation
		isValidName, isValidEmail, isValidTickets := validateUserInput(userInfo, remainingTickets)

		if isValidName && isValidEmail && isValidTickets {
			
			remainingTickets, bookings = bookTickets(remainingTickets, userInfo, bookings)
	
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

func getuserDetails() userData {
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

	var userInfo = userData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	return userInfo
}

func validateUserInput(userInfo userData, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(userInfo.firstName)>=2 && len(userInfo.lastName)>=2
	isValidEmail := strings.Contains(userInfo.email, "@")
	isValidTickets := userInfo.numberOfTickets > 0 && userInfo.numberOfTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTickets 
}

func bookTickets(remainingTickets uint, userInfo userData, bookings []userData) (uint, []userData) {
	remainingTickets = remainingTickets - userInfo.numberOfTickets
	bookings = append(bookings, userInfo)

	fmt.Printf("Thank you %s for booking %d tickets. You will recieve a confirmation email at %v \n", userInfo.firstName, userInfo.numberOfTickets, userInfo.email)
	fmt.Printf("%v tickets still remaining.\n", remainingTickets)

	return remainingTickets, bookings
}

func getFirstNames(bookings []userData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}