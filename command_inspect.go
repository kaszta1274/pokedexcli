package main

import "fmt"

func commandInspect(cfg *config, arguments []string) error {
	pokemonName := arguments[0]

	pokemon, inPokedex := cfg.pokedex[pokemonName]
	if !inPokedex {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, pokemonStat := range pokemon.Stats {
		fmt.Printf(" -%v: %v\n", pokemonStat.Stat.Name, pokemonStat.BaseStat)
	}

	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf(" -%v\n", pokemonType.Type.Name)
	}

	return nil
}
