package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) < 2 {
		return errors.New("no location provided")
	}

	locationAreaName := args[1]

	resp, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemons in %s:\n", resp.Name)
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf("\t- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
