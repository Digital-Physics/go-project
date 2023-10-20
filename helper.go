// this is mainly used as a demo to show us how to split our code over multiple files
// a little different than python; we have to put the package we want to have access in the helper file
package main

import "strings"

// package functions or variables you want to export must be Capitalized (although we didn't here and it still worked), just like the libraries we import like strings.Contains()
// function variables require types and then show return values. multiple return values are possible like Python.
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	// package functions you make names must be Capitalized, just like the libraries we import like strings.Contains()
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
