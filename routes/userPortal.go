package routes

import (
	"assignment-2/handler/admin"
	"assignment-2/handler/user"
	"assignment-2/util"
	"fmt"
)

func UserPortal() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	for {
		userPortalDisplay()
	x:
		var userInput string
		fmt.Println("Please select between 1 to 5:")

		if err := util.SingleLineInput(&userInput); err != nil {
			fmt.Println(err)
			goto x
		}
		if userInput == "!q" {
			return
		}
		i, err := util.InputParserStringToInt(&userInput)
		if err != nil {
			fmt.Println(err)
			goto x
		}
		switch i {
		case 1:
			admin.ViewAllVenues() //Show all venues => can reuse admin's view all venues
		case 2:
			user.BookVenue() // Add to bookingslist
		case 3:
			user.EditBooking() //Get booking and amend => need unique id
		case 4:
			user.CancelBooking() // remove booking from list
		case 5:
			return
		default:
			fmt.Printf("Invalid input: \"%s\".", userInput)
			goto x
		}

	}
}

func userPortalDisplay() {
	util.DisplayHeader("Welcome to the Venue Booking System")
	fmt.Println("1. View all venues")
	fmt.Println("2. Book Venue")
	fmt.Println("3. Edit Booking")
	fmt.Println("4. Cancel Booking")
	fmt.Println("5. Log out as User")
	fmt.Println()
}

// func createUser() *customer.Customer {

// 	// c := customer.New()
// 	return c
// }
