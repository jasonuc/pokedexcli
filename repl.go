package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startingRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedexcli >")

		scanner.Scan()

		text := scanner.Text()
		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}

		err := command.callback(cfg, cleaned...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Lists location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous location areas ",
			callback:    callbackMapB,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Lists the pokemons in a location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attempt to catch a pokemon and add it to your pokedex",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "Shows details on pokemon in your pokedex",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays pokemons in pokedex",
			callback:    callbackPokedex,
		},
	}
}

func cleanInput(s string) []string {
	lowered := strings.ToLower(s)
	words := strings.Fields(lowered)

	return words
}
