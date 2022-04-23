package main

import (
	"assignment-2/routes"
	"assignment-2/util"
	"fmt"
)

func main() {
	
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	for {
		mainDisplay()
	x:
		fmt.Println("Please select 1 or 2: ")
		var userInput string
		err := util.SingleLineInput(&userInput)
		if err != nil {
			fmt.Println(err)
			goto x
		}
		i, err := util.InputParserStringToInt(&userInput)
		if err != nil {
			fmt.Println(err)
			goto x
		}
		switch i {
		case 1:
			routes.UserPortal()
		case 2:
			routes.AdminPortal()
		default:
			fmt.Printf("Invalid input: \"%s\".\n", userInput)
			goto x
		}

	}
}

func mainDisplay() {
	util.DisplayHeader("Venue Booking System")
	fmt.Println("1. User Portal")
	fmt.Println("2. Admin Portal")
	fmt.Println()
}
