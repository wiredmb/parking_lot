package commands

import (
	"fmt"
	"parking"
	"perror"
	"ptypes"
	"strconv"
)

// CmdCreateParking ...
type CmdCreateParking struct {
	Command
	capacity ptypes.Capacity
}

// NewCommandCreateParking new create parking command instance
func NewCommandCreateParking(command string, parsedCommand []string) *CmdCreateParking {
	var cmd = new(CmdCreateParking)
	cmd.Cmd = command
	cmd.Args = parsedCommand[1:]
	return cmd
}

// Name returns command name
func (ccp *CmdCreateParking) Name() string {
	return ccp.Cmd
}

// Execute executes the create parking command
func (ccp *CmdCreateParking) Execute() error {
	var err error

	if err = ccp.verify(); err != nil {
		return err
	}

	if err = ccp.run(); err != nil {
		return err
	}

	return nil
}

func (ccp *CmdCreateParking) verify() error {
	if 1 != len(ccp.Args) {
		return perror.ErrInvalidParams
	}

	return nil
}

func (ccp *CmdCreateParking) run() error {
	val, err := strconv.ParseUint(ccp.Args[0], 0, 64)
	if nil != err {
		return perror.ErrInvalidParams
	}
	ccp.capacity = ptypes.Capacity(val)

	if 1 > ccp.capacity {
		return perror.ErrInvalidParams
	}

	station := parking.NewStation(ccp.capacity)
	parking.Set(station)
	ccp.Output = fmt.Sprintf("Created a parking lot with %v slots", ccp.capacity)

	return nil
}
