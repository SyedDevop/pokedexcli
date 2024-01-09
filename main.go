package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func getLine() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Pokedex > ")
	line, err := reader.ReadString('\n')
	if err != nil {
		if err != io.EOF {
			fmt.Fprintf(os.Stderr, "unable readline: %+v\n", err)
		}
		fmt.Printf("Thanks for playing!")
		os.Exit(0)
	}
	return strings.TrimSpace(line)
}

func main() {
	commands := getCommands()
	clint := NewClient()
	for {
		line := getLine()
		command, exists := commands[line]
		if !exists {
			fmt.Printf("Unknown command: {%s} This are the available commands: \n", line)
			commands["help"].callback(nil)
			continue
		}
		err := command.callback(clint)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
