package main

import (
	"fmt"
)

// CliCommand represents a command in the CLI application.
type CliCommand struct {
	callback    func() error
	name        string
	description string
}

// commandHelp is the callback for the 'help' command.
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

// GetCommands returns a map of available CLI commands.
func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
