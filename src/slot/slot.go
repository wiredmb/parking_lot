package slot

import (
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
