package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var Commands = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Get help on the Pokedex commands",
		callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}

func CleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	return strings.Fields(lowerCase)
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`)
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputWords := CleanInput(scanner.Text())
		command, ok := Commands[inputWords[0]]
		if ok {
			err := command.callback()
			fmt.Printf("Errors: %v\n", err)
		} else {
			fmt.Println("Unknown command")
		}
	}
}
