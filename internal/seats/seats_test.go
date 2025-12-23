package seats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAllocateSeatWithCustomSeatType tests allocating a seat with a custom SeatType passed dynamically in the test case
func TestAllocateSeatWithCustomSeatType(t *testing.T) {
	// Define a custom seat type dynamically for this test case (this is the "custom" type you want to test)
	customSeatType := SeatType("VIP") // Explicitly convert the string "VIP" to a SeatType

	// Define a custom seat pattern with a mix of standard and the custom seat type passed as a SeatType
	customPattern := []SeatType{
		customSeatType, LowerBerth, UpperBerth, customSeatType,
		SideLower, SideUpper, customSeatType, LowerBerth,
	}

	// Initialize seats with the custom pattern
	seatsList := InitializeSeats(9, customPattern) // Assume 3 bogeys

	// Try allocating a seat of the custom seat type ("VIP")
	preferredSeatType := customSeatType
	seat, err := AllocateSeat(seatsList, preferredSeatType)

	// Ensure that a seat of the custom type is allocated
	assert.NoError(t, err)
	assert.NotNil(t, seat)
	assert.Equal(t, preferredSeatType, seat.Type) // The allocated seat should be "VIP"

	// Now book the custom seats and try allocating the custom seat again
	for i := range seatsList {
		if seatsList[i].Type == customSeatType {
			seatsList[i].IsBooked = true
		}
	}

	// Try allocating a custom seat after booking all "VIP" seats
	seat, err = AllocateSeat(seatsList, preferredSeatType)

	// Ensure fallback logic works and a different seat type is allocated
	assert.NoError(t, err)
	assert.NotNil(t, seat)
	assert.NotEqual(t, preferredSeatType, seat.Type) // The allocated seat should not be "VIP"
}

// TestAllocateSeatPreferredUnavailable tests allocating a preferred seat when it's unavailable
func TestAllocateSeatPreferredUnavailable(t *testing.T) {
	// Initialize seats with default pattern (using 5 bogeys)
	seatsList := InitializeSeats(5, SeatPattern)

	// Book all "Lower Berth" seats to simulate them being unavailable
	for i := range seatsList {
		if seatsList[i].Type == LowerBerth {
			seatsList[i].IsBooked = true
		}
	}

	// Try allocating a preferred "Lower Berth" seat
	preferredSeatType := LowerBerth
	seat, err := AllocateSeat(seatsList, preferredSeatType)

	// Ensure that fallback logic is used, and a seat is allocated
	assert.NoError(t, err)
	assert.NotNil(t, seat)
	assert.NotEqual(t, preferredSeatType, seat.Type) // Should not be "LB"
}

// TestAllocateSeatAllAvailable tests allocating a preferred seat when all seats are available
func TestAllocateSeatAllAvailable(t *testing.T) {
	seatsList := InitializeSeats(2, SeatPattern) // 2 bogeys
	preferredSeatType := LowerBerth

	seat, err := AllocateSeat(seatsList, preferredSeatType)

	assert.NoError(t, err)
	assert.NotNil(t, seat)
	assert.Equal(t, preferredSeatType, seat.Type) // Should allocate the preferred seat
}

// TestAllocateSeatNoSeatsAvailable tests behavior when all seats are booked
func TestAllocateSeatNoSeatsAvailable(t *testing.T) {
	seatsList := InitializeSeats(2, SeatPattern)

	// Book all seats
	for i := range seatsList {
		seatsList[i].IsBooked = true
	}

	preferredSeatType := UpperBerth
	seat, err := AllocateSeat(seatsList, preferredSeatType)

	assert.Error(t, err)
	assert.Nil(t, seat)
}
