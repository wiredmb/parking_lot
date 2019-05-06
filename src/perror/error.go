package perror

import "errors"

var (
	// ErrParkingFull error for parking station is full
	ErrParkingFull = errors.New("Sorry, parking station is full")
	// ErrSlotNotAvailable error when slot is not available
	ErrSlotNotAvailable = errors.New("Slot not available")
	// ErrInvalidSlotNumber error for invalid slot number
	ErrInvalidSlotNumber = errors.New("Invalid slot number")
	// ErrNoEngagedSlots ...
	ErrNoEngagedSlots = errors.New("All slots are available")

	// Zero ...
	Zero = 0
)
