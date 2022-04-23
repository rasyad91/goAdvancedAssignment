package routes

import (
	"assignment-2/handler/admin"
	"assignment-2/util"
	"fmt"
)

func AdminPortal() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	for {
		adminPortalDisplay()
	x:
		fmt.Println("Please select between 1 to 6")
		var userInput string
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
			admin.ViewAllVenues() //done
		case 2:
			admin.ViewBookingsOfVenue() //done
		case 3:
			admin.AddVenue() //done
		case 4:
			admin.DeleteVenue() //done
		case 5:
			admin.EditVenue() //done
		case 6:
			return
		default:
			fmt.Printf("Invalid input: \"%s\"\n", userInput)
			goto x
		}

	}
}

func adminPortalDisplay() {
	util.DisplayHeader("Venue Booking System [Admin]")
	fmt.Println("1. View all venues")
	fmt.Println("2. View all bookings")
	fmt.Println("3. Add venue")
	fmt.Println("4. Delete venue")
	fmt.Println("5. Edit venue")
	fmt.Println("6. Log out as Admin")
	fmt.Println()
}
