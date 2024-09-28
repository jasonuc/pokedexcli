package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) < 2 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[1]

	pokemon, ok := cfg.caught[pokemonName]

	if !ok {
		return fmt.Errorf("haven't caught %s yet", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t-%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	for _, typ := range pokemon.Types {
		fmt.Printf("\t- %s\n", typ.Type.Name)
	}
	return nil
}
