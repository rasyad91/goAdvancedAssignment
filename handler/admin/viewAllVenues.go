package admin

import (
	"assignment-2/data/venue"
	"assignment-2/util"
	"fmt"
)

func ViewAllVenues() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	util.DisplayHeader("Venues")
	err := venue.ViewAll()
	if err != nil {
		fmt.Println(err)

	}
	for {
		var userInput string
		fmt.Println("Sort by: 1. Name | 2. Capacity | 3. Type | 4. Cost | 5. Location")
		fmt.Println("6. Return to Admin Menu")
		fmt.Print("Enter value: ")
		err := util.SingleLineInput(&userInput)
		if err != nil {
			fmt.Println(err)
			continue
		}
		i, err := util.InputParserStringToInt(&userInput)
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch i {
		case 1:
			fmt.Println("Sort by Name - Ascending order:")
			s := venue.SortByName()
			for _, v := range s {
				fmt.Printf("ID: %d, Name: %s, Type: %s, Cost: $%.2f/session, Capacity: %d, Location: %s, Description: \"%s\"\n", v.ID, v.Name, v.Type, v.CostPerSession, v.Capacity, v.Location, v.Description)
			}
		case 2:
			fmt.Println("Sort by Capacity - Descending order:")
			s := venue.SortByCapacity()
			for _, v := range s {
				fmt.Printf("ID: %d, Capacity: %d, Type: %s, Cost: $%.2f/session, Name: %s, Location: %s, Description: \"%s\"\n", v.ID, v.Capacity, v.Type, v.CostPerSession, v.Name, v.Location, v.Description)
			}
		case 3:
			fmt.Println("Sort by Type - Ascending order:")
			s := venue.SortByType()
			for _, v := range s {
				fmt.Printf("ID: %d,  Type: %s, Name: %s,Cost: $%.2f/session, Capacity: %d, Location: %s, Description: \"%s\"\n", v.ID, v.Type, v.Name, v.CostPerSession, v.Capacity, v.Location, v.Description)

			}
		case 4:
			fmt.Println("Sort by Cost Per Session - Descending order:")
			s := venue.SortByCost()
			for _, v := range s {
				fmt.Printf("ID: %d, Cost: $%.2f/session, Name: %s, Type: %s,  Capacity: %d, Location: %s, Description: \"%s\"\n", v.ID, v.CostPerSession, v.Name, v.Type, v.Capacity, v.Location, v.Description)

			}
		case 5:
			fmt.Println("Sort by Location - Ascending order:")
			s := venue.SortByLocation()
			for _, v := range s {
				fmt.Printf("ID: %d, Location: %s, Name: %s, Type: %s, Cost: $%.2f/session, Capacity: %d,  Description: \"%s\"\n", v.ID, v.Location, v.Name, v.Type, v.CostPerSession, v.Capacity, v.Description)

			}
		case 6:
			return
		default:
			fmt.Println("Please select between 1 to 6")

		}
	}

}
