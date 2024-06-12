package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("No pokemon name provided")
	}

	pokemonName := args[0]

	if _, ok := cfg.pokemons[pokemonName]; !ok {
		fmt.Printf("%s has not been caught!\n", pokemonName)
		return nil
	}

	pokemon := cfg.pokemons[pokemonName]

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("- %s\n", typ.Type.Name)
	}

	return nil
}
