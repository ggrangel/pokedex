package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	c := cfg.pokeapiClient

	if len(args) != 1 {
		return errors.New("No pokemon name provided")
	}

	pokemonName := args[0]

	if _, ok := cfg.pokemons[pokemonName]; ok {
		fmt.Printf("%s has already been caught!\n", pokemonName)
		return nil
	}

	pokemon, err := c.GetPokemon(pokemonName)

	if err != nil {
		return err
	}

	randNum := rand.Intn(pokemon.BaseExperience)
	const threshold = 50

	if randNum > threshold {
		fmt.Printf("Failed to catch %s!\n", pokemonName)
		return nil
	}

	// add the caught pokemon to the config
	cfg.pokemons[pokemonName] = pokemon

	fmt.Printf("%s was caught!\n", pokemonName)

	return nil
}
