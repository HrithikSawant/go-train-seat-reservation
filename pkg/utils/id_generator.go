package utils

import "fmt"

var bookingIDCounter = 1

// GenerateBookingID generates a unique sequential booking ID in the format "BOOK-001"
func GenerateBookingID() string {
	// Generate a booking ID and increment the counter for the next call
	bookingID := fmt.Sprintf("BOOK-%03d", bookingIDCounter)
	bookingIDCounter++ // Increment the counter
	return bookingID
}
