package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("Pokemon(s) in Pokedex:")

	for _, pokemon := range cfg.caught {
		fmt.Printf("\t- %s\n", pokemon.Name)
	}

	return nil
}
