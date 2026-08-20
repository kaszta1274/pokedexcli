package main

import "fmt"

func commandExplore(cfg *config, arguments []string) error {
	locationName := arguments[0]

	locationDetailsResp, err := cfg.client.GetLocationDetails(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationDetailsResp.PokemonEncounters {
		fmt.Printf(" - %v\n", pokemon.Pokemon.Name)
	}

	return nil
}
