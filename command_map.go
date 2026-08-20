package main

import (
	"fmt"
)

func commandMap(cfg *config, arguments []string) error {
	locationsResp, err := cfg.client.GetLocations(cfg.Next)
	if err != nil {
		return err
	}

	for _, area := range locationsResp.Results {
		fmt.Println(area.Name)
	}

	cfg.Previous = locationsResp.Previous
	cfg.Next = locationsResp.Next

	return nil
}

func commandMapBack(cfg *config, arguments []string) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locationsResp, err := cfg.client.GetLocations(cfg.Previous)
	if err != nil {
		return err
	}

	for _, area := range locationsResp.Results {
		fmt.Println(area.Name)
	}

	cfg.Previous = locationsResp.Previous
	cfg.Next = locationsResp.Next

	return nil
}
