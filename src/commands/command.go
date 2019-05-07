package commands

var (
	// CreateParkingLot ...
	CreateParkingLot = "create_parking_lot"
	// ParkVehicle ...
	ParkVehicle = "park"
	// RemoveVehicle ...
	RemoveVehicle = "leave"
	// GetParkingLotStatus ...
	GetParkingLotStatus = "status"
	// GetRegNumWithColor ...
	GetRegNumWithColor = "registration_numbers_for_cars_with_colour"
	// GetSlotNumWithColor ...
	GetSlotNumWithColor = "slot_numbers_for_cars_with_colour"
	// GetSlotNumWithRegNum ...
	GetSlotNumWithRegNum = "slot_number_for_registration_number"
	// ExitParkingLot ...
	ExitParkingLot = "exit"
)

// ICommand interface for all parking lot commands
type ICommand interface {
	Name() string
	Execute() error
}

// Command struct
type Command struct {
	Cmd    string
	Args   []string
	Output string
}
