package main

import "github.com/kaszta1274/pokedexcli/internal/pokeapi"

type config struct {
	commands map[string]cliCommand
	client   *pokeapi.Client
	Next     *string
	Previous *string
}
