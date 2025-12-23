package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"train-booking/internal/booking"
	"train-booking/internal/seats"
	utils "train-booking/pkg/utils"
)

func main() {
	// Initialize seats for SINGLE train
	seatsList := seats.InitializeSeats(9, seats.SeatPattern)
	reader := bufio.NewReader(os.Stdin)

	for {
		showMainMenu()

		fmt.Print("Enter choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			handleBooking(seatsList, reader)
		case "2":
			showSeatMatrix(seatsList, reader)
		case "3":
			fmt.Println("\nThank you for using Train Booking CLI 🚆")
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("\nInvalid choice. Please try again.")
		}
	}
}

func showMainMenu() {
	fmt.Println("========================================")
	fmt.Println("    TRAIN BOOKING SYSTEM - CLI")
	fmt.Println("========================================")
	fmt.Println("Train : SJ-101 - San Jose Express")
	fmt.Println("Route : SAN JOSE → SAN FRANCISCO")
	fmt.Println("----------------------------------------")
	fmt.Println("1. Book Ticket")
	fmt.Println("2. View Seat Matrix")
	fmt.Println("3. Exit")
	fmt.Println()
}

func handleBooking(seatsList []seats.Seat, reader *bufio.Reader) {
	fmt.Println("\n----------------------------------------")
	fmt.Println("            BOOK TICKET")
	fmt.Println("----------------------------------------")

	fmt.Print("Passenger Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("\nSeat Preference:")
	fmt.Println("1. Lower Berth (LB)")
	fmt.Println("2. Middle Berth (MB)")
	fmt.Println("3. Upper Berth (UB)")
	fmt.Println("4. Side Lower (SL)")
	fmt.Println("5. Side Upper (SU)")
	fmt.Println("6. No Preference")
	fmt.Print("Enter choice: ")

	prefInput, _ := reader.ReadString('\n')
	prefInput = strings.TrimSpace(prefInput)

	var preferredSeat seats.SeatType
	switch prefInput {
	case "1":
		preferredSeat = seats.LowerBerth
	case "2":
		preferredSeat = seats.MiddleBerth
	case "3":
		preferredSeat = seats.UpperBerth
	case "4":
		preferredSeat = seats.SideLower
	case "5":
		preferredSeat = seats.SideUpper
	default:
		preferredSeat = ""
	}

	bookingID := utils.GenerateBookingID()
	passengers := []booking.Passenger{
		{Name: name, PreferredSeat: preferredSeat},
	}

	b := booking.NewBooking(bookingID, passengers)

	seat, err := b.AllocateSeat(seatsList, preferredSeat)
	if err != nil {
		fmt.Println("\n❌ Train is FULL")
		fmt.Println("Booking Status : WAITLIST")
		waitForEnter(reader)
		return
	}

	b.Passengers[0].AllocatedSeat = seat

	fmt.Println("\n========================================")
	fmt.Println("          TICKET CONFIRMED 🎉")
	fmt.Println("========================================")
	fmt.Printf("Booking ID : %s\n", bookingID)
	fmt.Printf("Passenger  : %s\n", name)
	fmt.Printf("Bogey No   : %d\n", seat.BogeyNo)
	fmt.Printf("Seat No    : %d\n", seat.SeatNo)
	fmt.Printf("Berth Type : %s\n", seat.Type)
	fmt.Println("========================================")

	waitForEnter(reader)
}

func showSeatMatrix(seatsList []seats.Seat, reader *bufio.Reader) {
	fmt.Println("\n----------------------------------------")
	fmt.Println("           SEAT MATRIX VIEW")
	fmt.Println("----------------------------------------")

	// Header
	fmt.Print("SeatType : ")
	for _, st := range seats.SeatPattern {
		fmt.Printf("%-3s ", st)
	}
	fmt.Println()
	fmt.Println("----------------------------------------")

	seatsPerBogey := len(seats.SeatPattern)
	totalBogeys := len(seatsList) / seatsPerBogey

	for b := 1; b <= totalBogeys; b++ {
		fmt.Printf("Bogey %-2d : ", b)

		start := (b - 1) * seatsPerBogey
		end := start + seatsPerBogey

		for i := start; i < end; i++ {
			if seatsList[i].IsBooked {
				fmt.Print("1   ")
			} else {
				fmt.Print("0   ")
			}
		}
		fmt.Println()
	}

	fmt.Println("----------------------------------------")
	fmt.Println("Legend: 0 = AVAILABLE | 1 = BOOKED")

	waitForEnter(reader)
}

func waitForEnter(reader *bufio.Reader) {
	fmt.Println("\nPress Enter to return to Main Menu...")
	reader.ReadString('\n')
}

