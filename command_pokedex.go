package main

import "fmt"

func commandPokedex(cfg *config, arguments []string) error {
	fmt.Println("Your Pokedex:")

	for pokemonName := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemonName)
	}

	return nil
}
