package commands

import (
	"fmt"
	"parking"
	"perror"
	"strings"
)

// CmdGetSlotNumWithColor ...
type CmdGetSlotNumWithColor struct {
	Command
	color string
}

// NewCommandGetSlotNumWithColor new get slot number with color command instance
func NewCommandGetSlotNumWithColor(command string, parsedCommand []string) *CmdGetSlotNumWithColor {
	var cmd = new(CmdGetSlotNumWithColor)
	cmd.Cmd = command
	cmd.Args = parsedCommand[1:]
	return cmd
}

// Name returns command name
func (cd *CmdGetSlotNumWithColor) Name() string {
	return cd.Cmd
}

// Execute executes the GetSlotNumWithColor command
func (cd *CmdGetSlotNumWithColor) Execute() error {
	var err error

	if err = cd.verify(); err != nil {
		return err
	}

	if err = cd.run(); err != nil {
		return err
	}

	return nil
}

func (cd *CmdGetSlotNumWithColor) verify() error {
	if 1 != len(cd.Args) || cd.Args[0] == perror.Empty {
		return perror.ErrInvalidParams
	}

	return nil
}

func (cd *CmdGetSlotNumWithColor) run() error {
	cd.color = cd.Args[0]

	slots, err := parking.Get().GetSlotsBy("color", cd.color)
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
