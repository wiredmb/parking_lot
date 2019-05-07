package commands

import (
	"fmt"
	"parking"
	"perror"
	"vehicle"
)

// CmdPark ...
type CmdPark struct {
	Command
	vehicle *vehicle.Vehicle
}

// NewCommandPark new park command instance
func NewCommandPark(command string, parsedCommand []string) *CmdPark {
	var cmd = new(CmdPark)
	cmd.Cmd = command
	cmd.Args = parsedCommand[1:]
	return cmd
}

// Name returns command name
func (cp *CmdPark) Name() string {
	return cp.Cmd
}

// Execute executes the park command
func (cp *CmdPark) Execute() error {
	var err error

	if err = cp.verify(); err != nil {
		return err
	}

	if err = cp.run(); err != nil {
		return err
	}

	return nil
}

func (cp *CmdPark) verify() error {
	if 2 != len(cp.Args) {
		return perror.ErrInvalidParams
	}

	return nil
}

func (cp *CmdPark) run() error {
	cp.vehicle = vehicle.NewVehicle(cp.Args[0], cp.Args[1])

	slot, err := parking.Get().ParkVehicle(cp.vehicle)
	if err != nil {
		return err
	}

	cp.Output = fmt.Sprintf("Allocated slot number: %v", slot.GetSlotNumber())

	return nil
}
