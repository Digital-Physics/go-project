package main

// fmt is Format package and has Print function; https://pkg.go.dev/fmt
import "fmt"

func main() {
	// var are mutable; const are immutable
	// types can be inferred
	var conferenceName = "Golang Event"
	// var conferenceName string = "Golang Event"
	// conferenceName := "Golang Event"
	// there are many unsigned and signed ints and floats
	// uint8, uint16, uint32, uint64, int8... int64, float32, float64, complex64, complex128
	// (uint8 is 8 bits, capable of representing 0 to 255)
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	// alternative ways to insert variables and do line returns
	// %v is the standard format, but you can use others for formatting numbers, seeing the Type, etc. https://pkg.go.dev/fmt
	fmt.Printf("Welcome to our %v application \n", conferenceName)
	fmt.Println("We have", conferenceTickets, "tickets and", remainingTickets, "tickets are still available.")
	fmt.Println("Get your tickets here.")

	// define a type because we don't assign a value so the program can't infer it yet
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// code that get's the user's name similar to input in Python. the & gives a pointer to userName. C and C++ have explicit pointers too.
	// we don't want to pass the value of the userName; we want to pass the reference to where it is in memory so it can be assigned by the user.
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for purchasing %v tickets. \nAn email was sent to %v \n", firstName, lastName, userTickets, email)
	fmt.Println("There are now", remainingTickets, "left")

	// & can also be used to print out a memory location for the variable value
	// fmt.Println(&userName)

}
