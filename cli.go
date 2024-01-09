package main

import (
	"fmt"
	"os"
)

type CliCommand struct {
	callback    func(clint *Clint) error
	name        string
	description string
}

func commandHelp(_ *Clint) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Println(cmd.description)
	}
	return nil
}

func commandExit(_ *Clint) error {
	fmt.Printf("Thanks for playing!")
	os.Exit(0)
	return nil
}

func commandMap(clint *Clint) error {
	data, err := clint.GetPokeList()
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, area := range data.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb(clint *Clint) error {
	if clint.prevUri == nil {
		return fmt.Errorf("no previous uri found")
	}
	data, err := clint.GetPokePrevesList()
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, area := range data.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "help: Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit: Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "map: command reveals 20 Pokemon world location names. Each subsequent use of 'map' unveils the next 20 locations in a continuous sequence",
			callback:    commandMap,
		},

		"mapb": {
			name:        "mapb",
			description: "mapb: command in the Pokemon world lists 20 location areas. Repeating the command shows the previous 20 locations, allowing you to navigate backward.",
			callback:    commandMapb,
		},
	}
}
