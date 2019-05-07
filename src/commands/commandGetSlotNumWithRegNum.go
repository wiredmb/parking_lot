package commands

import (
	"fmt"
	"parking"
	"perror"
	"strings"
)

// CmdGetSlotNumWithRegNum ...
type CmdGetSlotNumWithRegNum struct {
	Command
	registrationNumber string
}

// NewCommandGetSlotNumWithRegNum new get slot number with registration number command instance
func NewCommandGetSlotNumWithRegNum(command string, parsedCommand []string) *CmdGetSlotNumWithRegNum {
	var cmd = new(CmdGetSlotNumWithRegNum)
	cmd.Cmd = command
	cmd.Args = parsedCommand[1:]
	return cmd
}

// Name returns command name
func (cd *CmdGetSlotNumWithRegNum) Name() string {
	return cd.Cmd
}

// Execute executes the GetSlotNumWithRegNum command
func (cd *CmdGetSlotNumWithRegNum) Execute() error {
	var err error

	if err = cd.verify(); err != nil {
		return err
	}

	if err = cd.run(); err != nil {
		return err
	}

	return nil
}

func (cd *CmdGetSlotNumWithRegNum) verify() error {
	if 1 != len(cd.Args) || cd.Args[0] == perror.Empty {
		return perror.ErrInvalidParams
	}

	return nil
}

func (cd *CmdGetSlotNumWithRegNum) run() error {
	cd.registrationNumber = cd.Args[0]

	slots, err := parking.Get().GetSlotsBy("regnum", cd.registrationNumber)
	if err != nil {
		return err
	}

	var outputList = []string{}
	for _, s := range slots {
		outputList = append(
			outputList,
			fmt.Sprintf("%v", s.GetSlotNumber()),
		)
	}
	cd.Output = strings.Join(outputList, perror.Comma)

	return nil
}
