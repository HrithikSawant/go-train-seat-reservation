package booking

import (
	"train-booking/internal/seats"
)

// Passenger represents a passenger's booking details
type Passenger struct {
	Name          string
	PreferredSeat seats.SeatType
	AllocatedSeat *seats.Seat
}

// Booking represents a booking with a list of passengers
type Booking struct {
	BookingID  string
	Passengers []Passenger
}

// NewBooking initializes a new booking
func NewBooking(bookingID string, passengers []Passenger) Booking {
	return Booking{
		BookingID:  bookingID,
		Passengers: passengers,
	}
}

// AllocateSeat allocates a seat to the passenger based on preference
// It calls the AllocateSeat method from the seats package to handle the seat allocation.
func (b *Booking) AllocateSeat(seatsList []seats.Seat, preferredSeatType seats.SeatType) (*seats.Seat, error) {
	// Try to allocate the preferred seat first
	seat, err := seats.AllocateSeat(seatsList, preferredSeatType)
	if err != nil {
		// If the preferred seat is not available, return the error
		return nil, err
	}
	return seat, nil
}
