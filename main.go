package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/term"
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

// TODO : Do some code cleanup using boot.dev example.
func main() {
	command := getCommands()
	for {
		line := getLine()
		fmt.Printf("line: %v\n", line)

		if line == "exit" {
			fmt.Printf("Thanks for playing!")
			break
		}

		command[line].callback()

	}
}
