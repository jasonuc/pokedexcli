package main

import (
	"time"

	"github.com/jasonuc/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationUrl *string
	prevLocationUrl *string
	caught          map[string]pokeapi.Pokemon
}

func main() {

	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caught:        make(map[string]pokeapi.Pokemon),
	}

	startingRepl(&cfg)
}
