package admin

import (
	"assignment-2/data/booking"
	"assignment-2/data/venue"
	"assignment-2/util"
	"fmt"
	"sync"
)

func DeleteVenue() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(2)

	util.DisplayHeader("Delete Venue")
	venue.ViewIDName()
	fmt.Println()
	fmt.Println("Please select the ID of venue, you wish to delete.")
	fmt.Println("To exit without deleting type: !q")
	for {
		var userInput string
		fmt.Print("Venue ID: ")
		err := util.SingleLineInput(&userInput)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if userInput == "!q" {
			return
		}

		i, err := util.InputParserStringToInt(&userInput)
		if err != nil {
			fmt.Println(err)
			continue
		}

		go func(i int) {
			if err := venue.RemoveByID(i); err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}(i)

		go func(i int) {
			booking.RemoveBookingList(i)
			if err := recover(); err != nil {
				fmt.Println("Add new venue: ", err)
			}
			wg.Done()
		}(i)

		wg.Wait()
		break
	}
}
