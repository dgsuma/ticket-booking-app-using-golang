package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference" // only for string data types can use :
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}  // slice is dynamic in length, despite array is fixed in length

	greetUsers(conferenceName, conferenceTickets, remainingTickets)	// function call

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets) // function call

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName  // array assignment
			bookings = append(bookings, firstName + " " + lastName) // slice assignment

			// fmt.Printf("The whole slice: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("Slice type: %T\n", bookings)
			// fmt.Printf("Slice length: %v\n", len(bookings))

			// fmt.Println(remainingTickets)
			// fmt.Println(&remainingTickets)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := getFirstNames(bookings)  // function call
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 { 
				// end program
				fmt.Println("Our convention is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}

	/*
	// switch statement example
	city := "London"

	switch city {
		case "New York":
			// execute code for booking New York conference tickets
		case "Singapore", "Hong Kong":
			// some code here for Singapore & Hong Kong
		case "London", "Berlin":
			// some code here for London & Berlin
		case "Mexico City":
			// execute code for booking Mexico City conference tickets
		default:
			fmt.Print("No valid city selected")
	}
	*/
}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{} // slice declaration
	for _, booking := range bookings { // _ is blank identifier (to ignore a variable you don't want to use)
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) { // return values (bool, bool, bool)
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber // A Go function can return multiple values, here we return 3 bool values
}

func getUserInput() (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Please enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}