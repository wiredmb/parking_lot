package commands

import (
	"fmt"
	"parking"
	"perror"
	"strings"
)

// CmdStatus ...
type CmdStatus struct {
	Command
}

// NewCommandStatus new status command instance
func NewCommandStatus(command string, parsedCommand []string) *CmdStatus {
	var cmd = new(CmdStatus)
	cmd.Cmd = command
	cmd.Args = parsedCommand[1:]
	return cmd
}

// Name returns command name
func (cs *CmdStatus) Name() string {
	return cs.Cmd
}

// Execute executes the status command
func (cs *CmdStatus) Execute() error {
	var err error

	if err = cs.run(); err != nil {
		return err
	}

	return nil
}

func (cs *CmdStatus) run() error {
	slots, err := parking.Get().GetEngagedSlots()
	if err != nil {
		return err
	}

	var outputList = []string{
		fmt.Sprintf("%-10s%-25s%-10s",
			"Slot No.",
			"Registration No",
			"Colour",
		),
		fmt.Sprintf("%-10v%-25v%-10v",
			"----------",
			"-------------------------",
			"----------",
		),
	}

	for _, s := range slots {
		v := s.GetVehicle()
		outputList = append(
			outputList,
			fmt.Sprintf(
				"%-10v%-25v%-10v",
				s.GetSlotNumber(),
				v.GetRegistrationNumber(),
				v.GetColor(),
			),
		)
	}

	cs.Output = strings.Join(outputList, perror.Newline)

	return nil
}
