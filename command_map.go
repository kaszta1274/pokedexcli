package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	locationAreas, err := cfg.client.GetMap(cfg.Next)
	if err != nil {
		return err
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	cfg.Previous = locationAreas.Previous
	cfg.Next = locationAreas.Next

	return nil
}

func commandMapBack(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locationAreas, err := cfg.client.GetMap(cfg.Previous)
	if err != nil {
		return err
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	cfg.Previous = locationAreas.Previous
	cfg.Next = locationAreas.Next

	return nil
}
