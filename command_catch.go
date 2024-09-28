package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) < 2 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[1]

	resp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(resp.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	cfg.caught[pokemonName] = resp
	fmt.Printf("%s was caught\n", resp.Name)
	return nil
}
