package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/KOTBCAnorax/pokedex/internal/pokeAPI"
	"github.com/KOTBCAnorax/pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(args ...string) error
}

var Cache = pokecache.NewCache(10)

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
	"map": {
		name:        "map",
		description: "Display next 20 areas in the Pokemon world",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Display previous 20 areas in the Pokemon world",
		callback:    commandMapb,
	},
	"config": {
		name:        "config",
		description: "View current config",
		callback:    commandConfig,
	},
	"cache": {
		name:        "cache",
		description: "Display current cahce state",
		callback:    commandCache,
	},
	"explore": {
		name:        "explore {area name}",
		description: "Get a list of all pokemons in the given area",
		callback:    commandExplore,
	},
}

func CleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	return strings.Fields(lowerCase)
}

func commandExit(args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(args ...string) error {
	fmt.Print(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: See next 20 areas of the Pokemon world
mapb: See previous 20 areas of the Pokemon world
config: See previous and next URLs
explore {area name}: get a list of all pooemon in the given area
`)
	return nil
}

func commandMap(args ...string) error {
	pokeAPI.AdvanceMap(Cache)
	return nil
}

func commandMapb(args ...string) error {
	pokeAPI.RetreatMap(Cache)
	return nil
}

func commandConfig(args ...string) error {
	fmt.Println(pokeAPI.Config.Prev)
	fmt.Println(pokeAPI.Config.Next)
	return nil
}

func commandCache(args ...string) error {
	Cache.Display()
	return nil
}

func commandExplore(args ...string) error {
	if len(args) == 0 {
		fmt.Println("No area name provided")
		return fmt.Errorf("No area name provided")
	}
	pokeAPI.GetPokemonsList(args[0], Cache)
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputWords := CleanInput(scanner.Text())
		command, ok := Commands[inputWords[0]]
		if ok {
			if len(inputWords) > 1 {
				command.callback(inputWords[1])
			} else {
				command.callback()
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
