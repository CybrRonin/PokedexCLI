package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("you need to supply a valid Pokémon name")
	}

	name := args[0]
	pokemonData, err := cfg.pokeapiClient.PokemonInfo(name)
	if err != nil {
		return err
	}
	odds := (1000 - pokemonData.BaseExperience) / 100
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	if rand.Intn(100) > odds {
		fmt.Printf("%s escaped!\n", name)
	} else {
		fmt.Printf("%s was caught!\n", name)
		cfg.pokedex[name] = pokemonData
	}
	return nil
}
