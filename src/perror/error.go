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
	// ErrInvalidParams ...
	ErrInvalidParams = errors.New("Invalid command params")
	// ErrInvalidCommand ...
	ErrInvalidCommand = errors.New("Invalid command")

	// Zero ...
	Zero = 0
)
