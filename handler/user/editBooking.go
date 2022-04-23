package user

import (
	"assignment-2/data/booking"
	"assignment-2/data/venue"
	"assignment-2/util"
	"fmt"
)

func EditBooking() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	util.DisplayHeader("Edit Booking")
	booking.DisplayAllBookings()
	var sessions []bool
	var bookings *booking.Bookings
	var editBooking booking.Booking

	var userInput string

	for {
		fmt.Print("Please select the venue ID of the booking you wish to edit: ")

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
		} else if len(b.BookedSessions) == 0 {
			fmt.Println("ID selected has no bookings. Please select ID from the list above with bookings.")
		} else {
			bookings = b
			break
		}
	}

	for {
		fmt.Print("Please select the date of booking you wish to edit [yyyy-mm-dd]: ")
		if err := util.SingleLineInput(&userInput); err != nil {
			fmt.Println(err)
			continue
		}
		if userInput == "!q" {
			fmt.Println()
			break
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
		if !ok {
			fmt.Println("Date entered does not have a booking!")
		} else {
			editBooking.Date = userInput
			sessions = s
			fmt.Println()
			break
		}
	}

	booking.PrintBookedSessions(sessions)
	for {
		fmt.Print("Please select the session you wish to edit: ")
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
		if sessions[i-1] == false {
			fmt.Println("Session not booked! Please select booked session to cancel.")
			continue
		}
		editBooking.Session = i
		break
	}

	go func(b *booking.Booking) {
		booking.RemoveBooking(bookings.VenueID, b)
	}(&editBooking)

	venue.ViewIDName()

	for {
		fmt.Print("New venue ID [to skip type !q]: ")
		if err := util.SingleLineInput(&userInput); err != nil {
			fmt.Println(err)
			continue
		}
		if userInput == "!q" {
			b, found := booking.GetBookingsFromVenue(bookings.VenueID)
			if !found {
				fmt.Println("ID not found in venue list. Please select ID from the list above.")
			} else {
				bookings = b
				break
			}
			break
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
		fmt.Print("Please select the new date you would like to book [yyyy-mm-dd][to skip type !q]: ")
		if err := util.SingleLineInput(&userInput); err != nil {
			fmt.Println(err)
			continue
		}
		if userInput == "!q" {
			s, ok := bookings.BookedSessions[editBooking.Date]
			if !ok {
				fmt.Println("in edit, user input for date")
				bookings.BookedSessions[editBooking.Date] = make([]bool, 24)
				s = bookings.BookedSessions[editBooking.Date]
				fmt.Println(bookings.BookedSessions[editBooking.Date])
			}
			sessions = s
			break
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
		editBooking.Date = userInput
		fmt.Println()
		break
	}

	booking.PrintAvailableSessions(sessions)
	for {
		fmt.Print("Please select the new session [to skip type !q]: ")
		if err := util.SingleLineInput(&userInput); err != nil {
			fmt.Println(err)
			continue
		}
		if userInput == "!q" {
			break
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
		editBooking.Session = i
		break
	}

	go func(bs *booking.Bookings, b *booking.Booking) {
		bookings.Enqueue(b)
		if err := recover(); err != nil {
			fmt.Println("Add new venue: ", err)
		}
	}(bookings, &editBooking)

	c := make(chan string)
	go func(id int) {
		if v, _ := venue.GetByID(id); v == nil {
			fmt.Println("ID not found")
		} else {
			c <- v.Name
			close(c)
		}
	}(bookings.VenueID)

	util.DisplayHeader("New Booking summary")
	fmt.Println("Venue: ", <-c)
	fmt.Println("Date: ", editBooking.Date)
	fmt.Println("Session: ", editBooking.Session)
	if err := util.SingleLineInput(&userInput); err != nil {
		fmt.Println(err)
	}
}
