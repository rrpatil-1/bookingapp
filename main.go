package main

import "fmt"

func main() {

	var conferencename = "GO Conference"
	const ConferenceTicket = 50
	var remainingTikets int = 50

	// array has fix length  and less dyanamic than other language
	//var bookings [50]string
	// slices is another data type in go which is abstraction of array.
	// but is more efficient and dynamic to add element and delete
	var bookings []string

	fmt.Printf("Hello welcome to %v booking application\n", conferencename)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", ConferenceTicket, remainingTikets)
	fmt.Printf("Get you ticket here to attend.\n")

	var FirstName string
	var LastName string
	var email string
	var UserTickets int
	fmt.Println("Enter your firstname: ")
	fmt.Scan(&FirstName)

	fmt.Println("Enter your LastName: ")
	fmt.Scan(&LastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter your tickets: ")
	fmt.Scan(&UserTickets)

	//update remaining tickets
	remainingTikets = remainingTikets - UserTickets

	//bookings[0] = FirstName + " " + LastName

	bookings = append(bookings, FirstName+" "+LastName)

	// fmt.Printf("whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Array Type: %T\n", bookings)
	// fmt.Printf("Array length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. you will received the confirmation mail at %v. \n", FirstName, LastName, UserTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTikets, conferencename)

	fmt.Printf("These are all our bookings: %v\n", bookings)

}
