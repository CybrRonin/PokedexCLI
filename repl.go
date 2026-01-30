package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/CybrRonin/PokedexCLI/internal/pokeapi"
)

const PROMPT = "Pokedex > "

func startRepl(cfg *config) {
	reader := bufio.NewScanner((os.Stdin))
	//cache := NewCache(30 * time.Second)
	for {
		fmt.Print(PROMPT)
		reader.Scan()

		input := reader.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}

		cmd := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[cmd]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas in the Pokémon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokémon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location name>",
			description: "Lists all the pokenmon that can be found in a given location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon name>",
			description: "attempt to catch a Pokémon and add it to your Pokédex",
			callback:    commandCatch,
		},
	}
}
