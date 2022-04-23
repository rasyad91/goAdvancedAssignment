package venue

import (
	"errors"
	"fmt"
	"strings"
)

// Venue struct
type Venue struct {
	ID             int
	Name           string
	Capacity       int
	Location       string
	Type           string
	Description    string
	CostPerSession float64
}

type List []*Venue

var list List

func init() {
	list = []*Venue{
		{1, "Undercove", 30, "Lavender", "Event Space", "The Undercove is a cosy shophouse unit .", 100.00},
		{2, "Common Ground", 100, "MacPherson", "Co-working Space", "", 300.50},
		{3, "Foundry", 20, "City Hall", "Seminar Room", "The Foundry boasts a spectacular view of the Skyline", 300},
		{4, "Tea House", 60, "Tanjong Pagar", "Meeting Room", "Beautifully designed meeting rooms", 200},
		{5, "Lunarly", 200, "Buona Vista", "Corporate Event ", "The venue is well-furnished with a suite of amenities", 1000},
	}
}

func GetList() *List {
	return &list
}

func incrementID() int {
	max := 0
	for _, venue := range list {
		if venue.ID > max {
			max = venue.ID
		}
	}
	return max + 1
}

// Get by ID, returns the venue and the index of the venue in the slice
func GetByID(ID int) (*Venue, int) {
	first, last := 0, len(list)-1

	for first <= last {
		mid := (first + last) / 2
		if list[mid].ID == ID {
			return list[mid], mid
		}
		if ID > list[mid].ID {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}
	return nil, -1
}

// Add new Venue to venue lists
func Add(v *Venue) {
	v.ID = incrementID()
	list = append(list, v)
}

// RemoveByID Venues from venue lists and sanitizes any bookings made
func RemoveByID(ID int) error {
	_, i := GetByID(ID)
	if i == -1 {
		return errors.New("ID not found in venue list.")
	}
	*(&list) = append(list[:i], list[i+1:]...)

	return nil
}

// Edit Venue in List
func EditBy(v *Venue, id int) {
	for i, venue := range list {
		if venue.ID == id {
			list[i] = v
		}
	}
}

// ViewAll venue lists
func ViewAll() error {
	if len(list) == 0 {
		return errors.New("No Venues in list. Please add venues to list")
	}
	for _, venue := range list {
		fmt.Println(venue)
	}
	return nil
}

// ViewSummary venue lists
func ViewIDName() error {
	if len(list) == 0 {
		return errors.New("No Venues in list. Please add venues to list")
	}
	for _, venue := range list {
		fmt.Printf("ID: %d, Venue: %s, Capacity: %d, Type: %s\n", venue.ID, venue.Name, venue.Capacity, venue.Type)
	}
	return nil
}

func (v *Venue) String() string {
	return fmt.Sprintf("ID: %d, Type: %s, Name: %s, Cost: $%.2f/session, Capacity: %d, Location: %s, Description: \"%s\"", v.ID, v.Type, v.Name, v.CostPerSession, v.Capacity, v.Location, v.Description)
}

// SearchByName searches list by Name parameter
func SearchByName(name *string) bool {

	for _, venue := range list {
		if strings.EqualFold(venue.Name, *name) {
			return true
		}
	}
	return false
}

func ViewSortCategory(s int) {

	t := make([]Venue, len(list))
	for i := range list {
		t[i] = *list[i]
	}

}

func mergeCapacity(s1, s2 []*Venue) []*Venue {

	p1, p2 := 0, 0
	result := []*Venue{}

	for p1 < len(s1) && p2 < len(s2) {
		if s1[p1].Capacity > s2[p2].Capacity {
			result = append(result, s1[p1])
			p1++
		} else {
			result = append(result, s2[p2])
			p2++
		}
	}

	if p1 == len(s1) {
		result = append(result, s2[p2:]...)
	} else {
		result = append(result, s1[p1:]...)
	}

	return result
}

func SortByCapacityHelper(slice []*Venue) []*Venue {

	if len(slice) == 1 {
		return slice
	}

	left := make(chan []*Venue)
	go func() {
		left <- SortByCapacityHelper(slice[:len(slice)/2])
	}()

	right := make(chan []*Venue)
	go func() {
		right <- SortByCapacityHelper(slice[len(slice)/2:])
	}()

	return mergeCapacity(<-left, <-right)
}

func SortByCapacity() []*Venue {
	l := list
	return SortByCapacityHelper(l)
}

func mergeName(s1, s2 []*Venue) []*Venue {

	p1, p2 := 0, 0
	result := []*Venue{}

	for p1 < len(s1) && p2 < len(s2) {
		if s1[p1].Name < s2[p2].Name {
			result = append(result, s1[p1])
			p1++
		} else {
			result = append(result, s2[p2])
			p2++
		}
	}

	if p1 == len(s1) {
		result = append(result, s2[p2:]...)
	} else {
		result = append(result, s1[p1:]...)
	}

	return result
}

// Sort implements mergesort algorithm
func SortByNameHelper(slice []*Venue) []*Venue {

	if len(slice) == 1 {
		return slice
	}

	left := make(chan []*Venue)
	go func() {
		left <- SortByNameHelper(slice[:len(slice)/2])
	}()

	right := make(chan []*Venue)
	go func() {
		right <- SortByNameHelper(slice[len(slice)/2:])
	}()

	return mergeName(<-left, <-right)
}

func SortByName() []*Venue {
	l := list
	return SortByNameHelper(l)
}

func mergeType(s1, s2 []*Venue) []*Venue {

	p1, p2 := 0, 0
	result := []*Venue{}

	for p1 < len(s1) && p2 < len(s2) {
		if s1[p1].Type < s2[p2].Type {
			result = append(result, s1[p1])
			p1++
		} else {
			result = append(result, s2[p2])
			p2++
		}
	}

	if p1 == len(s1) {
		result = append(result, s2[p2:]...)
	} else {
		result = append(result, s1[p1:]...)
	}

	return result
}

// Sort implements mergesort algorithm
func SortByTypeHelper(slice []*Venue) []*Venue {

	if len(slice) == 1 {
		return slice
	}

	left := make(chan []*Venue)
	go func() {
		left <- SortByTypeHelper(slice[:len(slice)/2])
	}()

	right := make(chan []*Venue)
	go func() {
		right <- SortByTypeHelper(slice[len(slice)/2:])
	}()

	return mergeType(<-left, <-right)
}

func SortByType() []*Venue {
	l := list
	return SortByTypeHelper(l)
}

func mergeLocation(s1, s2 []*Venue) []*Venue {

	p1, p2 := 0, 0
	result := []*Venue{}

	for p1 < len(s1) && p2 < len(s2) {
		if s1[p1].Location < s2[p2].Location {
			result = append(result, s1[p1])
			p1++
		} else {
			result = append(result, s2[p2])
			p2++
		}
	}

	if p1 == len(s1) {
		result = append(result, s2[p2:]...)
	} else {
		result = append(result, s1[p1:]...)
	}

	return result
}

// Sort implements mergesort algorithm
func SortByLocationHelper(slice []*Venue) []*Venue {

	if len(slice) == 1 {
		return slice
	}

	left := make(chan []*Venue)
	go func() {
		left <- SortByLocationHelper(slice[:len(slice)/2])
	}()

	right := make(chan []*Venue)
	go func() {
		right <- SortByLocationHelper(slice[len(slice)/2:])
	}()

	return mergeLocation(<-left, <-right)
}

func SortByLocation() []*Venue {
	l := list
	return SortByLocationHelper(l)
}

func mergeCost(s1, s2 []*Venue) []*Venue {

	p1, p2 := 0, 0
	result := []*Venue{}

	for p1 < len(s1) && p2 < len(s2) {
		if s1[p1].CostPerSession < s2[p2].CostPerSession {
			result = append(result, s1[p1])
			p1++
		} else {
			result = append(result, s2[p2])
			p2++
		}
	}

	if p1 == len(s1) {
		result = append(result, s2[p2:]...)
	} else {
		result = append(result, s1[p1:]...)
	}

	return result
}

// Sort implements mergesort algorithm
func SortByCostHelper(slice []*Venue) []*Venue {

	if len(slice) == 1 {
		return slice
	}

	left := make(chan []*Venue)
	go func() {
		left <- SortByCostHelper(slice[:len(slice)/2])
	}()

	right := make(chan []*Venue)
	go func() {
		right <- SortByCostHelper(slice[len(slice)/2:])
	}()

	return mergeCost(<-left, <-right)
}

func SortByCost() []*Venue {
	l := list
	return SortByCostHelper(l)
}
