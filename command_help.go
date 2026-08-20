package main

import (
	"fmt"
)

func commandHelp(cfg *config, arguments []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for _, command := range cfg.commands {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}
