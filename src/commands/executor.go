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
		case command == GetParkingLotStatus:
		case command == GetRegNumWithColor:
		case command == GetSlotNumWithColor:
		case command == GetSlotNumWithRegNum:
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
