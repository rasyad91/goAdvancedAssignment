package booking

import (
	"assignment-2/data/venue"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Booking struct holds the venue details and the customer that books the venue
// Session => 1 day segregated to 24 sessions, 1 session 1h
// Customer from customer package
// Next booking in line
type Booking struct {
	Date    string
	Session int
	next    *Booking
}

// var b []*Booking

// Bookings lists booking by Venue
// Uses linkedlist data structure, with Session added, based on session
type Bookings struct {
	VenueID        int
	front          *Booking
	size           int
	BookedSessions map[string][]bool // maps the date to the sessions
}

const dateLayout = "2006-01-02"

// map bookings to venue ID
var BookingsList = make(map[int]*Bookings)

func RemoveBooking(venueID int, r *Booking) error {
	b := BookingsList[venueID]
	if b.front == nil {
		return fmt.Errorf("Empty list, nothing to remove")
	}

	current := b.front

	if dateSessionParser(b.front) == dateSessionParser(r) {
		if b.front.next == nil {
			b.front = nil
			BookingsList[venueID] = &Bookings{}
		} else {
			b.front = current.next
		}
	} else {
		prev := b.front
		current = prev.next

		for current.next != nil {
			prev = current
			current = current.next
		}
		if dateSessionParser(current) == dateSessionParser(r) {
			prev.next = current.next
			current.next = nil
		}
	}

	b.BookedSessions[r.Date][r.Session-1] = false
	return nil
}

func DisplayAllBookings() {
	for i := 1; i <= len(BookingsList); i++ {
		v, _ := venue.GetByID(i)
		if v == nil {
			continue
		}

		fmt.Printf("Venue ID: %d | Venue Name: %s\n", i, v.Name)

		if BookingsList[i].front == nil {
			fmt.Println("No bookings made for this venue")
			fmt.Println()
			continue
		}

		currentNode := BookingsList[i].front
		for currentNode != nil {
			fmt.Printf("%v\n", currentNode)
			currentNode = currentNode.next
		}
		fmt.Println()
	}
}

func (v *Booking) String() string {
	return fmt.Sprintf("Date: %s | Session: %d", v.Date, v.Session)
}

func init() {
	list := venue.GetList()
	for _, venue := range *list {
		if _, ok := BookingsList[venue.ID]; !ok {
			BookingsList[venue.ID] = &Bookings{VenueID: venue.ID}
			BookingsList[venue.ID].BookedSessions = make(map[string][]bool, 24)
		}
	}

	booking1 := []*Booking{{Date: "2021-10-01", Session: 10}, {Date: "2021-10-01", Session: 11}, {Date: "2021-11-11", Session: 1}, {Date: "2021-11-11", Session: 4}, {Date: "2021-09-21", Session: 2}}
	for _, v := range booking1 {
		if _, ok := BookingsList[2].BookedSessions[v.Date]; !ok {
			BookingsList[2].BookedSessions[v.Date] = make([]bool, 24)
		}
		BookingsList[2].Enqueue(v)
	}
	booking2 := []*Booking{{Date: "2021-05-01", Session: 1}, {Date: "2021-05-07", Session: 21}, {Date: "2021-05-07", Session: 22}, {Date: "2021-05-07", Session: 23}, {Date: "2021-05-07", Session: 24}}
	for _, v := range booking2 {
		if _, ok := BookingsList[1].BookedSessions[v.Date]; !ok {
			BookingsList[1].BookedSessions[v.Date] = make([]bool, 24)
		}
		BookingsList[1].Enqueue(v)
	}

}

func InitializeBookingList(ID int) {
	BookingsList[ID] = &Bookings{VenueID: ID}
	BookingsList[ID].BookedSessions = make(map[string][]bool, 24)
}

func RemoveBookingList(ID int) {
	delete(BookingsList, ID)
}

func PrintBookingsList() {
	for k, v := range BookingsList {
		fmt.Println(k, v)
	}
}

func GetBookingsFromVenue(id int) (*Bookings, bool) {
	b, ok := BookingsList[id]
	return b, ok
}

func StringToDateParser(s *string) (time.Time, error) {

	d, err := time.Parse(dateLayout, *s)
	if err != nil {
		return time.Time{}, fmt.Errorf("\"%s\" Invalid date. Sample format: \"2020-10-01\"", *s)
	}
	if ok := time.Now().Before(d); !ok {
		return time.Time{}, fmt.Errorf("\"%s\" Invalid date. Date selected is backdated.", *s)
	}

	return d, nil
}

func PrintAllSessions(s []bool) {
	fmt.Println("All sessions: ")
	for i, s := range s {
		if i < 9 {
			if s == true {
				fmt.Printf("Session %d | 0%d00-0%d00: Booked\n", i+1, i, i+1)
			} else {
				fmt.Printf("Session %d | 0%d00-0%d00: Available\n", i+1, i, i+1)
			}
		} else if i == 9 {
			if s == true {
				fmt.Printf("Session %d | 0%d00-%d00: Booked\n", i+1, i, i+1)
			} else {
				fmt.Printf("Session %d | 0%d00-%d00: Available\n", i+1, i, i+1)
			}
		} else {
			if i == 23 {
				if s == true {
					fmt.Printf("Session %d | %d00-0000: Booked\n", i+1, i)
				} else {
					fmt.Printf("Session %d | %d00-0000: Available\n", i+1, i)
				}
			} else {
				if s == true {
					fmt.Printf("Session %d | %d00-%d00: Booked\n", i+1, i, i+1)
				} else {
					fmt.Printf("Session %d | %d00-%d00: Available\n", i+1, i, i+1)
				}
			}
		}

	}
	fmt.Println()
}

func PrintBookedSessions(s []bool) {
	fmt.Println("Booked sessions: ")
	for i, s := range s {
		if i < 9 {
			if s == true {
				fmt.Printf("Session %d | 0%d00-0%d00: Booked\n", i+1, i, i+1)
			}
		} else if i == 9 {
			if s == true {
				fmt.Printf("Session %d | 0%d00-%d00: Booked\n", i+1, i, i+1)
			}
		} else {
			if i == 23 {
				if s == true {
					fmt.Printf("Session %d | %d00-0000: Booked\n", i+1, i)
				}
			} else {
				if s == true {
					fmt.Printf("Session %d | %d00-%d00: Booked\n", i+1, i, i+1)
				}
			}
		}

	}
	fmt.Println()
}

func PrintAvailableSessions(s []bool) {
	fmt.Println("Available sessions: ")
	for i, s := range s {
		if i < 9 {
			if s == false {
				fmt.Printf("Session %d | 0%d00-0%d00: Available\n", i+1, i, i+1)
			}
		} else if i == 9 {
			if s == false {
				fmt.Printf("Session %d | 0%d00-%d00: Available\n", i+1, i, i+1)
			}
		} else {
			if i == 23 {
				if s == false {
					fmt.Printf("Session %d | %d00-0000: Available\n", i+1, i)
				}
			} else {
				if s == false {
					fmt.Printf("Session %d | %d00-%d00: Available\n", i+1, i, i+1)
				}
			}
		}

	}
	fmt.Println()
}

func (q *Bookings) Enqueue(newBooking *Booking) {
	if q.front == nil {
		q.front = newBooking
		q.BookedSessions[newBooking.Date][newBooking.Session-1] = true
		q.size++
		return
	}

	if dateSessionParser(newBooking) < dateSessionParser(q.front) {
		newBooking.next = q.front
		q.front = newBooking
		q.BookedSessions[newBooking.Date][newBooking.Session-1] = true
		q.size++

		return
	}

	current := q.front
	for current.next != nil && dateSessionParser(newBooking) > dateSessionParser(current.next) {
		current = current.next
	}

	if current.next == nil {
		current.next = newBooking
	} else {
		newBooking.next = current.next
		current.next = newBooking
	}
	q.BookedSessions[newBooking.Date][newBooking.Session-1] = true
	q.size++
}

func dateSessionParser(b *Booking) int {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("DateSessionParser: ", err)
		}
	}()

	s := b.Date
	slice := strings.Split(s, "-")
	bt := []byte(strconv.Itoa(b.Session - 1))
	if len(bt) == 1 {
		bt = append([]byte{'0'}, bt...)
	}
	st := string(bt)

	slice = append(slice, st)
	joined := strings.Join(slice, "")
	i, _ := strconv.Atoi(joined)
	return i
}

func (q *Bookings) PrintAll() {
	fmt.Println("In print all")
	if q.front == nil {
		fmt.Println("Nothing to print! Empty list!")
		return
	}

	currentNode := q.front

	for currentNode != nil {
		if currentNode == q.front {
			fmt.Print("[Front] ")
		}
		fmt.Printf("%v ->", *currentNode)
		currentNode = currentNode.next
	}
	fmt.Println("")
}
