package admin

import (
	"assignment-2/data/booking"
	"assignment-2/data/venue"
	"assignment-2/util"
	"errors"
	"fmt"
)

func AddVenue() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	util.DisplayHeader("Add Venue")
	fmt.Println("Please key in the details of the venue.")

	var userInput string
	var newVenue venue.Venue

	for {
		if err := venueName(&userInput, ""); err != nil {
			fmt.Println(err)
			continue
		}
		util.StringTitle(&userInput)
		newVenue.Name = userInput
		break
	}

	for {
		i, err := venueCapacity(&userInput, "")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if i < 1 {
			fmt.Printf("Invalid capacity: \"%d\", Capactiy must be more than 0.\n", i)
		} else {
			newVenue.Capacity = i
			break
		}
	}

	for {
		if err := venueLocation(&userInput, ""); err != nil {
			fmt.Println(err)
			continue
		}
		util.StringTitle(&userInput)
		newVenue.Location = userInput
		break
	}

	for {
		if err := venueType(&userInput, ""); err != nil {
			fmt.Println(err)
			continue
		}
		util.StringTitle(&userInput)
		newVenue.Type = userInput
		break
	}

	for {
		f, err := venueCostPerSession(&userInput, "")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if f < 0.01 {
			fmt.Printf("Invalid cost per session: \"%.2f\", Capactiy must be more than 0.\n", f)
		} else {
			newVenue.CostPerSession = f
			break
		}
	}

	for {
		fmt.Println("Venue description: ")
		if err := util.MultiLinesInput(&userInput); err != nil {
			fmt.Println(err)
			continue
		}
		newVenue.Description = userInput
		break
	}
	c := make(chan int, 1)
	go func(v *venue.Venue) {
		venue.Add(v)
		if err := recover(); err != nil {
			fmt.Println("Add new venue: ", err)
		}
		c <- v.ID
		close(c)
	}(&newVenue)

	go func() {
		booking.InitializeBookingList(<-c)
		if err := recover(); err != nil {
			fmt.Println("Add new venue: ", err)
		}
	}()

	fmt.Println("New venue added!")
	fmt.Println(&newVenue)
	fmt.Scanln()
}

// Helper functions

func venueName(userInput *string, message string) error {
	fmt.Print("Venue name: ")
	if len(message) != 0 {
		fmt.Println(message)
	}
	if err := util.SingleLineInput(userInput); err != nil {
		return err
	}
	if err := util.InputNotEmpty(userInput); err != nil {
		return err
	}
	if venue.SearchByName(userInput) == true {
		return errors.New(fmt.Sprintf("\"%s\" already exist. Please use another name.\n", *userInput))
	}
	return nil
}

func venueCapacity(userInput *string, message string) (int, error) {
	fmt.Print("Venue capacity: ")
	if len(message) != 0 {
		fmt.Println(message)
	}
	if err := util.SingleLineInput(userInput); err != nil {
		return -1, err
	}
	i, err := util.InputParserStringToInt(userInput)
	if err != nil {
		return -1, err
	}

	return i, nil
}

func venueLocation(userInput *string, message string) error {
	fmt.Print("Venue location: ")
	if len(message) != 0 {
		fmt.Println(message)
	}
	if err := util.SingleLineInput(userInput); err != nil {
		return err
	}
	if err := util.InputNotEmpty(userInput); err != nil {
		return err
	}
	return nil
}

func venueType(userInput *string, message string) error {
	fmt.Print("Venue type: ")
	if len(message) != 0 {
		fmt.Println(message)
	}
	if err := util.SingleLineInput(userInput); err != nil {
		return err
	}
	if err := util.InputNotEmpty(userInput); err != nil {
		return err
	}
	return nil
}

func venueCostPerSession(userInput *string, message string) (float64, error) {
	fmt.Print("Venue cost per session: ")
	if len(message) != 0 {
		fmt.Println(message)
	}
	if err := util.SingleLineInput(userInput); err != nil {
		return -1, err
	}
	f, err := util.InputParserStringToFloat64(userInput)
	if err != nil {
		return -1, err
	}
	return f, nil
}
