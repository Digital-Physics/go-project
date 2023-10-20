// all our code must belong to a package
// the first statement in Go must be "package ..."; main is the standare main package name
package main

// fmt is Format package and has Print functions; https://pkg.go.dev/fmt
import (
	"fmt"
	"sync"
	"time"
	// "go-project/helper"
)

// var are mutable; const are immutable
const conferenceTickets int = 50

// types can be inferred
// var conferenceName string = "Golang Event"
// conferenceName := "Golang Event"
// := syntax should not be used for Package-level variables (like global vars)
var conferenceName = "Golang Event"

// maps are like dictionaries/hash tables but are not as nice as python's.
// you have to have the same types so we convert them to strings in this case
// this is like a list of empty hash tables where key will be string and value will be string.
// var bookings = make([]map[string]string, 0)
// structs should be used instead
// structs are like class objects but without the methods
// seems like people can and like to line up the types when defining them
// this will be used in one of our functions
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// there are many unsigned and signed ints and floats
// uint8, uint16, uint32, uint64, int8... int64, float32, float64, complex64, complex128
// (uint8 is 8 bits, capable of representing 0 to 255)
var remainingTickets uint = 50

// define the size of the array and the type of the array. no type mixing within the array.
// you can initialize some strings to start if you want: var bookings = [50]string{"Alice", "Bob"}
// alternative syntax for empty: var bookings = [50]string{}
// var bookings [50]string
// similar to Python indexing. can skip if you wanted as shown below.
// bookings[0] = "Alice"
// bookings[10] = "Bob"
// but a "slice", which is defined as a variable length array (like a python list), is a better solution. no length defined ahead of time.
//
// this makes a list of UserData objects with 0 of them to start.
// it will be in scope for a function to append to
var bookings = make([]UserData, 0)

// a WaitGroup allows us to do concurrency. it will hold the goroutine threads and wait for them to finish before exiting
var wg = sync.WaitGroup{}

// entry point
func main() {

	greetUsers()

	// there are no while loops. this is equivalent to a while True
	for {
		firstName, lastName, email, userTickets := getUserInput()
		// isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets)
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		// demo of switching as an alternative to if-else if-else. it's like case statements
		// city := "London"

		// switch city {
		// 	case "New York":
		// 		// execute code
		// 	case "Singapore", "Hong Kong":
		// 		// execute code
		// 	case "London":
		// 		// execute code
		// 	default:
		// 		fmt.Println("No valid city selected")
		// }

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			// we're adding 1 new thread to the wait group so the program doesn't end before this is returned
			// if we had a second go routing right below it, we'd say wg.Add(2)
			// we are increasing the counter used to track completed threads by 1
			wg.Add(1)
			// this go call is how we call a goroutine on the new thread
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our event is fully booked.")
				// we'll keep looping until this point
				// note: we'll need to book the exact number of remaining to get this program to end
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
	// we'll wait until all of our sleeping/email sending threads are done before ending the function
	wg.Wait()
}

func greetUsers() {
	// alternative ways to insert variables and do line returns. like Python fstrings and .format()
	// %v is the standard format, but you can use others for formatting numbers, seeing the Type, etc. https://pkg.go.dev/fmt
	// fmt.Printf("Welcome to our %v booking application \n", conferenceName)
	// fmt.Println("We have", conferenceTickets, "tickets and", remainingTickets, "tickets are still available.")
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// the type is a list, or slice, of strings. it can grow in length and unlike an array the len/memory doesn't neede to be defined/allocated in advanced
// note: the return value does not need {}
func getFirstNames() []string {
	firstNames := []string{} // initialized empty list
	// for loop notes:
	// break and continue can be used in for loops like they are in Python
	// there is no while loop in Go, only a for loop. so while true is just for. and while condition is for condition.
	// range allows us to iterate over elements. this is like enumerate() in Python, but the index is not used (i.e. _ indicates this)
	// _ makes the unused index value not be highlighted because the intellisense knows
	for _, booking := range bookings {
		// this splits the string using white space as a separator
		// var names = strings.Fields(booking)
		// and then we take the first name and append it
		// firstNames = append(firstNames, names[0])
		// go uses append like Python, but you pass the list and what you want to append into a function
		// and structs use . notation for getting property ("attribute" in python) names
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// if you return multiple values they need to go in (). are these considered "tuples" in Go?
func getUserInput() (string, string, string, uint) {
	// define the types of values you will get from the user input
	// define a type because we can't assign a value so the program can't infer it yet
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// code that get's the user's name similar to input() in Python. the & gives a pointer to userName. C and C++ have explicit pointers too.
	// we don't want to pass the value (make a copy of the firstName); we want to pass the reference to where it is in memory so the value can be assigned by the user.
	// & can also be used to print out a memory location for the variable value
	// fmt.Println(&userName)
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// put the input types in. this function returns no value so no extra syntax up front
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	// access those package (almost like global) variables that are in scope
	remainingTickets = remainingTickets - userTickets

	// initialize a UserData struct which was defined above and is in scope
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// append to package-level list that is in scope
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are now %v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// sleep is similar to python, but I don't remember the unit being needed in python
	// this is done as a surrogate for a long running process like generating a pdf and sending an email
	time.Sleep(50 * time.Second)
	// like an fstring in python, we use this to create a string and instert formatted variables
	var ticket = fmt.Sprintf("There were %v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("Sending tickets:\n %v \nsent to email address %v\n", ticket, email)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// this let's the wait group know this process is done so it doesn't need to wait any longer
	// we are decreasing the counter used to track completed threads by 1
	wg.Done()
}

// to create the go.mod file which initializes our project/module run "go mod init <project path>" or "go mod init go-project" in this case
// the module path can correspond to where you plan to publish your module on github
// the go.mod file that is created has the name of the project and the version of go
//
// to run the file: "go run main.go"
// to run them all (notice there is no import of the helper.go module from main) "go run main.go helper.go" or "go run ." to run all
