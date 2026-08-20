package main

import (
	"github.com/kaszta1274/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := config{
		commands: getCommands(),
		client:   pokeapi.NewClient(),
		pokedex:  make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
