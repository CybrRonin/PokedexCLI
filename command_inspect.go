package main

import (
	"errors"
	"fmt"

	"github.com/CybrRonin/PokedexCLI/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("valid Pokémon name required")
	}

	name := args[0]
	pokemon, caught := cfg.pokedex[name]
	if !caught {
		return errors.New("No entry found for: " + name)
	}
	pokemonStats(pokemon)
	return nil
}

func pokemonStats(pokemon pokeapi.Pokemon) {
	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  - ", typeInfo.Type.Name)
	}
}
