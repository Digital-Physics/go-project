package main

// fmt is Format package and has Print function; https://pkg.go.dev/fmt
import (
	"fmt"
	"strings"
)

// entry point
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
	fmt.Printf("Welcome to our %v booking application \n", conferenceName)
	fmt.Println("We have", conferenceTickets, "tickets and", remainingTickets, "tickets are still available.")
	fmt.Println("Get your tickets here to attend.")

	// define the size of the array and the type of the array. no type mixing within the array.
	// you can initialize some strings to start if you want: var bookings = [50]string{"Alice", "Bob"}
	// alternative syntax for empty: var bookings = [50]string{}
	// var bookings [50]string
	// similar to Python indexing. could skip if you wanted as shown below.
	// bookings[0] = "Alice"
	// bookings[10] = "Bob"
	// but a "slice", which is defined as a variable length array (like a python list?), is a better solution. no length defined ahead of time.
	var bookings []string

	// there is no while loop in Go, only a for loop. so while true is just for. and while condition is for condition.
	for remainingTickets > 0 {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
		// isValidCity := city == "Singapore" || city == "London"
		// isInvalidCity := city != "Singapore" && city != "London"

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			// append is similar to python, but not a method on the list. it's a function.
			bookings = append(bookings, firstName+" "+lastName)
			// fmt.Printf("The whole slice: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("Slice type: %T\n", bookings)
			// fmt.Printf("Slice length: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for purchasing %v tickets. \nAn email was sent to %v \n", firstName, lastName, userTickets, email)
			fmt.Println("There are now", remainingTickets, "left")

			firstNames := []string{}
			// range allows us to iterate over elements
			// this is like for loop with enumerate() in Python, but the index is not used
			// _ makes the unused index value not be highlighted because the intellisense knows
			for _, booking := range bookings {
				// this splits the string using white space as a separator
				var names = strings.Fields(booking)
				// and then we take the first name and append it
				firstNames = append(firstNames, names[0])
			}
			fmt.Println("Bookings (first name only shown):", firstNames)

			// var noTicketsRemaining bool = remainingTickets == 0
			// noTicketsRemaining := remainingTickets == 0
			// if noTicketsRemaining {
			if remainingTickets == 0 {
				fmt.Println("Our conference is fully booked.")
				break // break loop, like Python
			}

			// & can also be used to print out a memory location for the variable value
			// fmt.Println(&userName)
			// } else if userTickets == remainingTickets {
			// 	// do something else
		} else {
			if !isValidName {
				fmt.Println("Your first name or last name is too short.")
			}
			if !isValidEmail {
				fmt.Println("Your email does not contain an @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Your ticket number is invalid")
			}
			// we use if - else, but break and continue work like Python in for loops
			// fmt.Printf("There are only %v tickets remaining, so you can't book %v tickets \n", remainingTickets, userTickets)
			// fmt.Println("Invalid data input")
		}
	}

}
