package admin

import (
	"assignment-2/data/venue"
	"assignment-2/util"
	"fmt"
)

func EditVenue() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	util.DisplayHeader("Edit Venue")
	venue.ViewIDName()

	var userInput string
	var existingVenue *venue.Venue
	var id int

	for {
		i, err := editVenueInput(&userInput)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if userInput == "!q" {
			return
		}
		existingVenue, id = venue.GetByID(i)
		if existingVenue == nil {
			fmt.Println("ID not found. Please select ID from the table above.")
			continue
		}
		break
	}

	for {
		fmt.Println(existingVenue)
		err := venueName(&userInput, "To skip type: !s")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if userInput == "!q" {
			return
		}
		if userInput != "!s" {
			util.StringTitle(&userInput)
			existingVenue.Name = userInput
		}
		break
	}

	for {
		fmt.Println(existingVenue)
		i, err := venueCapacity(&userInput, "To skip type: -1")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if i != -1 {
			existingVenue.Capacity = i
		}
		break
	}

	for {
		fmt.Println(existingVenue)
		err := venueLocation(&userInput, "To skip type: !s")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if userInput != "!s" {
			util.StringTitle(&userInput)
			existingVenue.Location = userInput
		}
		break

	}

	for {
		fmt.Println(existingVenue)
		err := venueType(&userInput, "To skip type: !s")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if userInput != "!s" {
			util.StringTitle(&userInput)
			existingVenue.Type = userInput
		}
		break

	}

	for {
		fmt.Println(existingVenue)
		f, err := venueCostPerSession(&userInput, "To skip type: -1")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if f != -1 {
			existingVenue.CostPerSession = f
		}
		break

	}

	for {
		fmt.Println("Venue description: ")
		err := util.MultiLinesInput(&userInput)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if userInput != "" {
			util.StringTitle(&userInput)
			existingVenue.Description = userInput
		}
		break
	}

	venue.EditBy(existingVenue, id)

}

// Helper functions

func editVenueInput(userInput *string) (int, error) {
	fmt.Println("Please select the ID of the Venue you wish to edit.")
	fmt.Print("Enter ID: ")
	err := util.SingleLineInput(userInput)
	if err != nil {
		return -1, err
	}
	i, err := util.InputParserStringToInt(userInput)
	if err != nil {
		return -1, err
	}
	return i, nil
}
