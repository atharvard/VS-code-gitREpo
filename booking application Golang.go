package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

//THEES are the PACKAGE FUNCTION  can assecces inside any function
var conferenceName = "Go Conference"
var conferenceTickets = 50 // here data tpye is not define why so? becauese golang consider data type if values are mention imediatly
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // creating empty map with initial size zero

// array comes in role for multiple users size 50 i.e for 50 bookings
// type creats a new type, with the name you specify(created a type called unseData based on struct)
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers() // function is called

	for { // infinite loop
		firstName, lastName, email, userTickets := getUserInput()

		//validation point i.e check point
		isValidName, isValidEmail, isValidTicketsNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		//isValidCity := city == "Singapore" || city == "London"
		//!isValidCity // ! results of above statements
		//isInvalidCity := city != "Singapore" && city != "London" // for invalid condition

		if isValidName && isValidEmail && isValidTicketsNumber {
			bookTicket(userTickets, firstName, lastName, email)
			go sendTickets(userTickets, firstName, lastName, email)

			//here call the function print first name
			firstNames := getFirstNames()
			fmt.Printf("These are all the first names of bookings: %v\n", firstNames) //cltr + c to move out of terminal

			//var noTicketsRemaining bool =  remainingTickets == 0 // these may be wrritten as(noTicketsRemaining := remainingTickets = 0)but we using or once , so need of these
			if remainingTickets == 0 {
				//end the program
				fmt.Println("Our conference is booked out. See you soon! ")
				break //it will terminate loop instantly
			}

		} else {

			//fmt.Printf("we only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			if !isValidName {
				fmt.Printf("First name nad Last name is too short\n")
			}
			if !isValidEmail {
				fmt.Printf("Email address you enter is not correct\n")

			}
			if !isValidTicketsNumber {
				fmt.Printf("Number of tickets you enter is Invalid\n")
			}

		}

	}

}

func greetUsers() {
	fmt.Printf("welcome to your %v booking application\n", conferenceName) // these will not execute untill we call it by its name
	fmt.Printf("we have total of %vtickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend") // we done this in printf so %v may be different for different output

}

func getFirstNames() []string { //inside input, outside output karan we return to the main function

	firstNames := []string{}           // below is for extracting omnly the first name from the bookings
	for _, booking := range bookings { // range gives iteration provieds the index and value for each element,  _ ignore thevvariable which we don`t want to use`
		//var names = strings.Fields(booking) //string.Fields -- split the string with space separtor i.e string=atharva deshpande into slice "atharva" "deshpande"

		firstNames = append(firstNames, booking.firstName)
	} // these loop ends when iterated over all elements of booking list then going to the nextr statement of our program
	return firstNames // thes can return the data as output in the main function
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for theirr name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) // for users personal info i.e asking for user input. Here by using pointers values are store in the variables

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName) // use of ppointers

	fmt.Println("Enter your Email: ")
	fmt.Scan(&email)

	fmt.Println("Enter No. of tickets you want: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//Creating a  map for a user
	//var userData = make(map[string]string) // make will create an empty map (string string -- for keyname and key value)

	//struct these will give user data objects
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	//userData["firstName"] = firstName      //assigning key name and value respectively
	//userData["lastName"] = lastName
	//userData["email"] = email

	// cannot mix data type in map ,in strct no need of coversion
	//userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) //formatUint is for coversion other data type to string part of strcov package, base10

	bookings = append(bookings, userData) // here slice comes into picture
	fmt.Printf("list of bookings is %v \n", bookings)

	//fmt.Print("The whole slice: %v\n", bookings)
	//fmt.Print("The first value: %v\n", bookings[0])
	//fmt.Print("The slice type: %T\n", bookings)
	//fmt.Print("slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a conformation Email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v,Hurry up! Book tickets now\n", remainingTickets, conferenceName)

}
func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName) // will save thes tickets in ticket sprint to save it in spring variable
	fmt.Println("#########")
	fmt.Printf("sending ticket\n %v \n to email addresss %v\n", ticket, email)
	fmt.Printf("########")
}
