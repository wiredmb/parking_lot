package parking

import (
	"parking_lot/src/ptypes"
	"parking_lot/src/slot"
)

// Station is a parking station
type Station struct {
	capacity ptypes.Capacity
	slots    []*slot.Slot
}

// NewStation returns new parking station
func NewStation(capacity ptypes.Capacity) *Station {
	station := new(Station)
	station.init(capacity)
	return station
}

func (ps *Station) init(capacity ptypes.Capacity) {
	ps.capacity = capacity
	ps.slots = make([]*slot.Slot, capacity)
	for idx := range ps.slots {
		ps.slots[idx] = slot.NewSlot(idx)
	}
}
