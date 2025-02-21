package main

import (
	"bufio"
	"fmt"
	"math/rand"
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
	"catch": {
		name:        "catch",
		description: "Try to catch given pokemon",
		callback:    commandCatch,
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
explore {area name}: Get a list of all pooemon in the given area
catch {pokemon name}: Try to catch given pokemon
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
		fmt.Println("no area name provided")
		return fmt.Errorf("no area name provided")
	}
	err := pokeAPI.GetPokemonsList(args[0], Cache)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func commandCatch(args ...string) error {
	if len(args) == 0 {
		fmt.Println("No pokemon name given")
		return fmt.Errorf("no pokemon name given")
	}

	pokemon, err := pokeAPI.GetPokemon(args[0], Cache)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	fmt.Printf("%s's base experience: %d\n", pokemon.Name, pokemon.BaseExperience)
	chance := int(100.0 - 100*float64(pokemon.BaseExperience)/maxBaseExp)
	fmt.Printf("Chance of success: %d%%\n", chance)
	randInt := rand.Intn(101)
	if randInt >= chance {
		Pokedex[pokemon.Name] = *pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}

var Pokedex = map[string]pokeAPI.Pokemon{}
var maxBaseExp = 390.0

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
