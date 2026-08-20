package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, arguments []string) error {
	pokemonName := arguments[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemonResp, err := cfg.client.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	if checkIfCaught(pokemonResp.BaseExperience) {
		fmt.Printf("%s was caught!\n", pokemonResp.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.pokedex[pokemonResp.Name] = pokemonResp
	} else {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
	}

	return nil
}

func checkIfCaught(BaseExperience int) bool {
	n := rand.Intn(BaseExperience)
	return n < 50
}
