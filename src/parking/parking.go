package parking

import (
	"perror"
	"ptypes"
	"slot"
	"vehicle"
)

// Station is a parking station
type Station struct {
	capacity     ptypes.Capacity
	startIndex   ptypes.Index
	engagedCount ptypes.Index
	slots        []*slot.Slot
}

var parkingStation *Station

// NewStation returns new parking station
func NewStation(capacity ptypes.Capacity) *Station {
	station := new(Station)
	station.init(capacity)
	return station
}

func (ps *Station) init(capacity ptypes.Capacity) {
	ps.capacity = capacity
	ps.startIndex = 1
	ps.engagedCount = 0
	ps.slots = make([]*slot.Slot, capacity)
	for idx := range ps.slots {
		ps.slots[idx] = slot.NewSlot(ptypes.Index(idx) + 1)
	}
}

// Get returns the parking station.
func Get() *Station {
	if parkingStation == nil {
		panic("Parking station is not initialized")
	}
	return parkingStation
}

// Set sets the parking station
func Set(station *Station) {
	parkingStation = station
}

// ParkVehicle parks the vehicle in available slot, else returns error.
func (ps *Station) ParkVehicle(vehicle *vehicle.Vehicle) (*slot.Slot, error) {
	availableSlot, err := ps.getAvailableSlot()
	if err != nil {
		return availableSlot, err
	}

	err = availableSlot.Allot(vehicle)
	if err != nil {
		return availableSlot, err
	}

	ps.engagedCount++
	return availableSlot, nil
}

func (ps *Station) getAvailableSlot() (*slot.Slot, error) {
	for _, slot := range ps.slots {
		if slot.IsAvailable() {
			return slot, nil
		}
	}
	return nil, perror.ErrParkingFull
}

// RemoveVehicleBySlotNumber removes the vehicle from parking slot by slot number.
func (ps *Station) RemoveVehicleBySlotNumber(slotNumber ptypes.Index) error {
	slot, err := ps.GetSlot(slotNumber)
	if nil != err {
		return err
	}

	slot.Free()
	ps.engagedCount--
	return nil
}

// GetSlot get the slot by slot number, else returns error
func (ps *Station) GetSlot(slotNumber ptypes.Index) (*slot.Slot, error) {
	if slotNumber < ps.startIndex && slotNumber > ptypes.Index(ps.capacity) {
		return nil, perror.ErrInvalidSlotNumber
	}
	return ps.slots[slotNumber-ps.startIndex], nil
}

// GetEngagedSlots returns slice of engaged parking slots
func (ps *Station) GetEngagedSlots() ([]*slot.Slot, error) {
	if ps.engagedCount == ptypes.Index(perror.Zero) {
		return nil, perror.ErrNoEngagedSlots
	}

	var slotIndex ptypes.Index
	engagedSlots := make([]*slot.Slot, ps.engagedCount)

	for _, slot := range ps.slots {
		if !slot.IsAvailable() {
			engagedSlots[slotIndex] = slot
			slotIndex++
		}
	}
	return engagedSlots, nil
}

// GetSlotsBy returns engaged slots with input vehicle color
func (ps *Station) GetSlotsBy(key, value string) ([]*slot.Slot, error) {
	var slots = make([]*slot.Slot, 0)
	var expected string
	for _, slot := range ps.slots {
		v := slot.GetVehicle()

		if nil == v {
			continue
		}

		switch key {
		case "regnum":
			expected = v.GetRegistrationNumber()
			break
		case "color":
			expected = v.GetColor()
			break
		}

		if value == expected {
			slots = append(slots, slot)
		}
	}

	return slots, nil
}
