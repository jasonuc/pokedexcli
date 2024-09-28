package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationUrl)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf("\t- %s\n", area.Name)
	}

	cfg.nextLocationUrl = resp.Next
	cfg.prevLocationUrl = resp.Previous

	return nil
}

func callbackMapB(cfg *config, args ...string) error {
	if cfg.prevLocationUrl == nil {
		return errors.New("you are still on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationUrl)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf("\t- %s\n", area.Name)
	}

	cfg.nextLocationUrl = resp.Next
	cfg.prevLocationUrl = resp.Previous

	return nil
}
