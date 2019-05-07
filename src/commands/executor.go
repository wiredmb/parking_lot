package commands

import (
	"bufio"
	"fmt"
	"log"
	"perror"
	"runtime"
	"strings"
)

// ExecuteCommands executes all commands.
func ExecuteCommands(scanner *bufio.Scanner) {
	newline := getNewlineStr()
	exit := false

	for !exit && scanner.Scan() {
		commandStr := strings.TrimRight(scanner.Text(), newline)
		parsedCommand := parse(commandStr)

		if len(parsedCommand) < 1 {
			log.Fatal("Insufficient command line params")
			return
		}

		command := parsedCommand[0]

		switch {
		case command == CreateParkingLot:
			cmdCreateParking := NewCommandCreateParking(command, parsedCommand)
			err := cmdCreateParking.Execute()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(cmdCreateParking.Output)

		case command == ParkVehicle:
			cmdPark := NewCommandPark(command, parsedCommand)
			err := cmdPark.Execute()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(cmdPark.Output)

		case command == RemoveVehicle:
			cmdRemove := NewCommandRemove(command, parsedCommand)
			err := cmdRemove.Execute()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(cmdRemove.Output)

		case command == GetParkingLotStatus:
			cmdStatus := NewCommandStatus(command, parsedCommand)
			err := cmdStatus.Execute()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(cmdStatus.Output)

		case command == GetRegNumWithColor:
			cmdGetRegNumWithColor := NewCommandGetRegNumWithColor(command, parsedCommand)
			err := cmdGetRegNumWithColor.Execute()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(cmdGetRegNumWithColor.Output)

		case command == GetSlotNumWithColor:
			cmdGetSlotNumWithColor := NewCommandGetSlotNumWithColor(command, parsedCommand)
			err := cmdGetSlotNumWithColor.Execute()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(cmdGetSlotNumWithColor.Output)

		case command == GetSlotNumWithRegNum:
			cmdGetSlotNumWithRegNum := NewCommandGetSlotNumWithRegNum(command, parsedCommand)
			err := cmdGetSlotNumWithRegNum.Execute()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(cmdGetSlotNumWithRegNum.Output)

		case command == ExitParkingLot && len(parsedCommand) == 1:
			exit = true
		default:
			fmt.Println(perror.ErrInvalidCommand)
		}
	}
}

func getNewlineStr() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}

func parse(input string) []string {
	s := strings.Split(input, " ")
	return s
}
