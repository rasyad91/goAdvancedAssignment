package admin

import (
	"assignment-2/data/booking"
	"assignment-2/util"
	"fmt"
)

func ViewBookingsOfVenue() {
	util.DisplayHeader("View All Bookings")
	booking.DisplayAllBookings()
	var u string
	if err := util.SingleLineInput(&u); err != nil {
		fmt.Println(err)
	}
}
