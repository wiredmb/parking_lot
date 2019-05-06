package slot

import (
	"parking_lot/src/perror"
	"parking_lot/src/ptypes"
	"parking_lot/src/vehicle"
)

// Slot is a parking slot with slot number and parked vehicle information.
type Slot struct {
	slotNumber ptypes.Index
	vehicle    *vehicle.Vehicle
}

// NewSlot return new parking slot assigning slot number.
func NewSlot(slotNumber ptypes.Index) *Slot {
	slot := new(Slot)
	slot.init(slotNumber)
	return slot
}

func (s *Slot) init(slotNumber ptypes.Index) {
	s.slotNumber = slotNumber
	s.vehicle = nil
}

// IsAvailable checks for availability of parking slot.
func (s *Slot) IsAvailable() bool {
	return s.vehicle == nil
}

// Allot allocates a vehicle into a parking slot.
func (s *Slot) Allot(vehicle *vehicle.Vehicle) error {
	if nil != s.vehicle {
		return perror.ErrSlotNotAvailable
	}

	s.vehicle = vehicle
	return nil
}

// Free frees up the parking slot.
func (s *Slot) Free() {
	s.vehicle = nil
}

// GetVehicle returns vehicle parked in slot.
func (s *Slot) GetVehicle() *vehicle.Vehicle {
	return s.vehicle
}
