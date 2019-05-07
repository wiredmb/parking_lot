package commands

import (
	"fmt"
	"parking"
	"perror"
	"ptypes"
	"strconv"
)

// CmdRemove ...
type CmdRemove struct {
	Command
	slotNumber ptypes.Index
}

// NewCommandRemove new remove command instance
func NewCommandRemove(command string, parsedCommand []string) *CmdRemove {
	var cmd = new(CmdRemove)
	cmd.Cmd = command
	cmd.Args = parsedCommand[1:]
	return cmd
}

// Name returns command name
func (cr *CmdRemove) Name() string {
	return cr.Cmd
}

// Execute executes the remove command
func (cr *CmdRemove) Execute() error {
	var err error

	if err = cr.verify(); err != nil {
		return err
	}

	if err = cr.run(); err != nil {
		return err
	}

	return nil
}

func (cr *CmdRemove) verify() error {
	if 1 != len(cr.Args) {
		return perror.ErrInvalidParams
	}

	return nil
}

func (cr *CmdRemove) run() error {
	var err error

	val, err := strconv.ParseUint(cr.Args[0], 0, 64)
	if nil != err {
		return perror.ErrInvalidParams
	}
	cr.slotNumber = ptypes.Index(val)

	if 1 > cr.slotNumber {
		return perror.ErrInvalidParams
	}

	err = parking.Get().RemoveVehicleBySlotNumber(cr.slotNumber)
	if err != nil {
		return err
	}

	cr.Output = fmt.Sprintf("Slot number %v is free", cr.slotNumber)

	return nil
}
