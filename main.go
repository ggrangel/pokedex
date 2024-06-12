package main

import (
	"time"

	"github.com/ggrangel/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaUrl *string
	prevLocationAreaUrl *string
	pokemons            map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Minute),
		pokemons:      make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}
