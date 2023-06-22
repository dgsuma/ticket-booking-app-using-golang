package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) { // return values (bool, bool, bool)
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber // A Go function can return multiple values, here we return 3 bool values
}