package seats

import (
	"errors"
)

// SeatType represents the type of a seat (e.g., LB, MB, UB, etc.)
type SeatType string

const (
	LowerBerth  SeatType = "LB"
	MiddleBerth SeatType = "MB"
	UpperBerth  SeatType = "UB"
	SideLower   SeatType = "SL"
	SideUpper   SeatType = "SU"
)

// Seat represents a seat in the train
type Seat struct {
	SeatNo   int
	BogeyNo  int
	Type     SeatType
	IsBooked bool
}

// SeatPattern holds the seat type pattern for each bogey
var SeatPattern = []SeatType{
	LowerBerth, MiddleBerth, UpperBerth, LowerBerth,
	MiddleBerth, UpperBerth, SideLower, SideUpper,
}

// CalculateSeatsPerBogey dynamically calculates the number of seats per bogey based on the pattern length
var seatsPerBogey = len(SeatPattern)

// CalculateTotalSeats calculates the total number of seats based on the number of bogeys
func CalculateTotalSeats(totalBogeys int) int {
	return seatsPerBogey * totalBogeys
}

// InitializeSeats creates and initializes a list of seats for the train
func InitializeSeats(totalBogeys int, customPattern []SeatType) []Seat {

	// Use default pattern if customPattern is nil
	if customPattern != nil {
		SeatPattern = customPattern
	}

	// Calculate total seats dynamically
	totalSeats := CalculateTotalSeats(totalBogeys)

	seats := make([]Seat, 0, totalSeats)

	// Create seats for each bogey
	seatNo := 1
	bogeyNo := 1
	for seatNo <= totalSeats {
		for _, seatType := range SeatPattern {
			seats = append(seats, Seat{
				SeatNo:   seatNo,
				BogeyNo:  bogeyNo,
				Type:     seatType,
				IsBooked: false,
			})
			seatNo++
		}
		bogeyNo++
	}

	return seats
}

// AllocateSeat allocates a seat based on the passenger's preference
// It first tries to allocate a preferred seat, and if unavailable, it allocates any available seat.
func AllocateSeat(seats []Seat, preferredSeatType SeatType) (*Seat, error) {
	// Try to allocate a preferred seat first
	preferredSeat, err := allocatePreferredSeat(seats, preferredSeatType)
	if err == nil {
		return preferredSeat, nil
	}

	// If no preferred seat is available, allocate any available seat, ensuring it's not the same type as the preferred one
	return allocateAnyAvailableSeat(seats, preferredSeatType)
}

// allocatePreferredSeat attempts to allocate a seat matching the preferred seat type
func allocatePreferredSeat(seats []Seat, preferredSeatType SeatType) (*Seat, error) {
	for i := range seats {
		if !seats[i].IsBooked && seats[i].Type == preferredSeatType {
			seats[i].IsBooked = true
			return &seats[i], nil
		}
	}
	return nil, errors.New("preferred seat type is not available")
}

// allocateAnyAvailableSeat allocates any available seat if no preferred seat is found
// It avoids allocating the same type as the preferred one
func allocateAnyAvailableSeat(seats []Seat, preferredSeatType SeatType) (*Seat, error) {
	for i := range seats {
		if !seats[i].IsBooked && seats[i].Type != preferredSeatType {
			seats[i].IsBooked = true
			return &seats[i], nil
		}
	}
	return nil, errors.New("no seats available")
}
