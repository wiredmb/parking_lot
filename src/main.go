package main

import (
	"bufio"
	"commands"
	"io"
	"log"
	"os"
)

var commandsInteractive io.Reader = os.Stdin

func main() {
	var scanner *bufio.Scanner
	argsLen := len(os.Args)

	switch {
	case argsLen == 2:
		commandsFile, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer commandsFile.Close()
		scanner = bufio.NewScanner(commandsFile)
	case argsLen == 1:
		scanner = bufio.NewScanner(commandsInteractive)
	default:
		log.Fatal("Unknown command line input")
	}

	commands.ExecuteCommands(scanner)
}
