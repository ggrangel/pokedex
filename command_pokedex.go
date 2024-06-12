package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {

	fmt.Println("Pokedex:")
	for _, pokemon := range cfg.pokemons {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}
