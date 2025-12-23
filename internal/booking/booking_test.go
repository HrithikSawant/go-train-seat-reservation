package booking

import (
	"testing"
	"train-booking/internal/seats"
	"train-booking/pkg/utils"

	"github.com/stretchr/testify/assert"
)

// TestBookingSeatAllocation tests that passengers are allocated seats based on their preferences.
func TestBookingSeatAllocation(t *testing.T) {
	// Initialize seats
	seatsList := seats.InitializeSeats(9, seats.SeatPattern)

	// Create passengers with preferences
	passengers := []Passenger{
		{Name: "Alice", PreferredSeat: seats.LowerBerth},
		{Name: "Bob", PreferredSeat: seats.UpperBerth},
	}

	// Create booking with passengers
	bookingID := "BOOK-001"
	booking := NewBooking(bookingID, passengers)

	// Allocate seats for passengers
	for i := range booking.Passengers {
		seat, err := booking.AllocateSeat(seatsList, booking.Passengers[i].PreferredSeat)
		// Ensure seat allocation succeeds
		assert.NoError(t, err)
		assert.NotNil(t, seat)
		booking.Passengers[i].AllocatedSeat = seat
	}

	// Verify seat allocation
	assert.Equal(t, "Alice", booking.Passengers[0].Name)
	assert.Equal(t, seats.LowerBerth, booking.Passengers[0].AllocatedSeat.Type)

	assert.Equal(t, "Bob", booking.Passengers[1].Name)
	assert.Equal(t, seats.UpperBerth, booking.Passengers[1].AllocatedSeat.Type)
}

// TestBookingIDGeneration tests the generation of a unique booking ID.
func TestBookingIDGeneration(t *testing.T) {
	// Generate two booking IDs
	id1 := utils.GenerateBookingID()
	id2 := utils.GenerateBookingID()

	// Ensure that the IDs are unique
	assert.NotEqual(t, id1, id2)

	// Ensure the format is "BOOK-XXX"
	assert.Regexp(t, "^BOOK-\\d{3}$", id1)
	assert.Regexp(t, "^BOOK-\\d{3}$", id2)
}

// TestBookingFailure tests the case where no seats are available to allocate.
func TestBookingFailure(t *testing.T) {
	// Initialize seats and book all seats
	seatsList := seats.InitializeSeats(9, seats.SeatPattern)
	for i := range seatsList {
		seatsList[i].IsBooked = true // Mark all seats as booked
	}

	// Create passengers with preferences
	passengers := []Passenger{
		{Name: "Alice", PreferredSeat: seats.LowerBerth},
		{Name: "Bob", PreferredSeat: seats.UpperBerth},
	}

	// Create booking with passengers
	bookingID := "BOOK-001"
	booking := NewBooking(bookingID, passengers)

	// Try allocating seats for passengers
	for i := range booking.Passengers {
		seat, err := booking.AllocateSeat(seatsList, booking.Passengers[i].PreferredSeat)

		// Assert that no seat is allocated
		assert.Error(t, err)
		assert.Nil(t, seat)
		booking.Passengers[i].AllocatedSeat = seat
	}

	// Verify that no seat was allocated to passengers
	assert.Nil(t, booking.Passengers[0].AllocatedSeat)
	assert.Nil(t, booking.Passengers[1].AllocatedSeat)
}
