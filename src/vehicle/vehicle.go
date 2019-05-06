package vehicle

// Vehicle represents a vehicle with registration number and color
type Vehicle struct {
	registrationNumber, color string
}

// NewVehicle return new vehicle assigning registration number and color of vehicle.
func NewVehicle(registrationNumber, color string) *Vehicle {
	vehicle := new(Vehicle)
	vehicle.init(registrationNumber, color)
	return vehicle
}

func (v *Vehicle) init(registrationNumber, color string) {
	v.registrationNumber = registrationNumber
	v.color = color
}

// GetRegistrationNumber returns registration number of vehicle.
func (v *Vehicle) GetRegistrationNumber() string {
	return v.registrationNumber
}

// GetColor returns color of vehicle.
func (v *Vehicle) GetColor() string {
	return v.color
}
