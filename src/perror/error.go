package perror

import "errors"

var (
	// ErrParkingFull error for parking station is full
	ErrParkingFull = errors.New("Sorry, parking lot is full")
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
	// ErrNotFound ...
	ErrNotFound = errors.New("Not found")

	// Zero ...
	Zero = 0
	// Newline ...
	Newline = "\n"
	// Empty ...
	Empty = ""
	// Comma ...
	Comma = ", "
)
