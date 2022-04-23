package user

import (
	"assignment-2/data/booking"
	"assignment-2/util"
	"fmt"
)

func CancelBooking() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	util.DisplayHeader("Cancel Booking")
	booking.DisplayAllBookings()
	var sessions []bool
	var bookings *booking.Bookings
	var removedBooking booking.Booking

	var userInput string

	for {
		fmt.Print("Please select the venue ID of the booking you wish to cancel: ")

		if err := util.SingleLineInput(&userInput); err != nil {
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
		fmt.Print("Please select the date of booking you wish to cancel [yyyy-mm-dd]: ")
		if err := util.SingleLineInput(&userInput); err != nil {
			fmt.Println(err)
			continue
		}
		if userInput == "!q" {
			return
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
			removedBooking.Date = userInput
			sessions = s
			fmt.Println()
			break
		}
	}

	booking.PrintBookedSessions(sessions)
	for {
		fmt.Print("Please select the session you wish to cancel: ")
		if err := util.SingleLineInput(&userInput); err != nil {
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
		if i > 24 || i < 1 {
			fmt.Println("Index out of range. Please select 1 - 24")
			continue
		}
		if sessions[i-1] == false {
			fmt.Println("Session not booked! Please select booked session to cancel.")
			continue
		}
		removedBooking.Session = i
		break
	}
	booking.RemoveBooking(bookings.VenueID, &removedBooking)
	util.DisplayHeader("Booking cancelled")
	fmt.Println("Venue: ", bookings.VenueID)
	fmt.Println("Date: ", removedBooking.Date)
	fmt.Println("Session: ", removedBooking.Session)

	if err := util.SingleLineInput(&userInput); err != nil {
		fmt.Println(err)
	}
}
