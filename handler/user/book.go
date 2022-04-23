package user

import (
	"assignment-2/data/booking"
	"assignment-2/data/venue"
	"assignment-2/util"
	"fmt"
)

func BookVenue() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover from Book Venue:", err)
		}
	}()

	var userInput string
	var bookings *booking.Bookings
	var newBooking booking.Booking
	var sessions []bool

	util.DisplayHeader("Book A Venue")
	venue.ViewIDName()

	fmt.Println("Please key in the details of the booking")

	for {
		fmt.Print("Venue ID: ")
		if err := util.SingleLineInput(&userInput); err != nil {
			fmt.Println(err)
			continue
		}
		i, err := util.InputParserStringToInt(&userInput)
		if err != nil {
			fmt.Println(err)
			continue
		}
		b, found := booking.GetBookingsFromVenue(i)
		if !found {
			fmt.Println("ID not found in venue list. Please select ID from the list above.")
		} else {
			bookings = b
			break
		}
	}

	for {
		fmt.Print("Please select the date you would like to book [yyyy-mm-dd]: ")
		if err := util.SingleLineInput(&userInput); err != nil {
			fmt.Println(err)
			continue
		}
		if err := util.InputNotEmpty(&userInput); err != nil {
			fmt.Println(err)
			continue
		}
		_, err := booking.StringToDateParser(&userInput)
		if err != nil {
			fmt.Println(err)
			continue
		}
		s, ok := bookings.BookedSessions[userInput]
		if ok && !util.Contains(s, false) {
			fmt.Println("The venue for this date is fully booked! Please select another date.")
			continue
		}
		if !ok {
			bookings.BookedSessions[userInput] = make([]bool, 24)
			s = bookings.BookedSessions[userInput]
		}
		sessions = s
		newBooking.Date = userInput
		fmt.Println()
		break
	}

	booking.PrintAvailableSessions(sessions)
	for {
		fmt.Print("Please select the session: ")
		if err := util.SingleLineInput(&userInput); err != nil {
			fmt.Println(err)
			continue
		}
		i, err := util.InputParserStringToInt(&userInput)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if i > 24 || i < 1 {
			fmt.Println("Index out of range. Please select 1 - 24")
			continue
		}
		if sessions[i-1] == true {
			fmt.Println("Session booked! Please select another session.")
			continue
		}
		newBooking.Session = i
		break
	}

	go func(b *booking.Booking) {
		bookings.Enqueue(b)
		if err := recover(); err != nil {
			fmt.Println("Add new venue: ", err)
		}
	}(&newBooking)

	c := make(chan string)
	go func(id int) {
		if v, _ := venue.GetByID(id); v == nil {
			fmt.Println("ID not found")
		} else {
			c <- v.Name
			close(c)
		}
	}(bookings.VenueID)

	util.DisplayHeader("Booking summary")
	fmt.Println("Venue: ", <-c)
	fmt.Println("Date: ", newBooking.Date)
	fmt.Println("Session: ", newBooking.Session)
	if err := util.SingleLineInput(&userInput); err != nil {
		fmt.Println(err)
	}
}

// Date string
// Session  int
// Next     *Booking
