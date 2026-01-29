package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("explore requires a location")
	}

	name := args[0]
	detailsResp, err := cfg.pokeapiClient.LocationDetails(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", detailsResp.Name)
	fmt.Println("Found Pokemon: ")
	for _, encounter := range detailsResp.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
