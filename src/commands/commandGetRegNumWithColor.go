package commands

import (
	"fmt"
	"parking"
	"perror"
	"strings"
)

// CmdGetRegNumWithColor ...
type CmdGetRegNumWithColor struct {
	Command
	color string
}

// NewCommandGetRegNumWithColor new get reg number with color command instance
func NewCommandGetRegNumWithColor(command string, parsedCommand []string) *CmdGetRegNumWithColor {
	var cmd = new(CmdGetRegNumWithColor)
	cmd.Cmd = command
	cmd.Args = parsedCommand[1:]
	return cmd
}

// Name returns command name
func (cd *CmdGetRegNumWithColor) Name() string {
	return cd.Cmd
}

// Execute executes the GetRegNumWithColor command
func (cd *CmdGetRegNumWithColor) Execute() error {
	var err error

	if err = cd.verify(); err != nil {
		return err
	}

	if err = cd.run(); err != nil {
		return err
	}

	return nil
}

func (cd *CmdGetRegNumWithColor) verify() error {
	if 1 != len(cd.Args) || cd.Args[0] == perror.Empty {
		return perror.ErrInvalidParams
	}

	return nil
}

func (cd *CmdGetRegNumWithColor) run() error {
	cd.color = cd.Args[0]

	slots, err := parking.Get().GetSlotsBy("color", cd.color)
	if err != nil {
		return err
	}

	var outputList = []string{}
	for _, s := range slots {
		v := s.GetVehicle()
		outputList = append(
			outputList,
			fmt.Sprintf("%v", v.GetRegistrationNumber()),
		)
	}
	cd.Output = strings.Join(outputList, perror.Comma)

	return nil
}
