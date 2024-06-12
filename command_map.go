package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	c := cfg.pokeapiClient

	resp, err := c.ListLocationAreas(cfg.nextLocationAreaUrl)

	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, v := range resp.Results {
		fmt.Printf(" - %s\n", v.Name)
	}

	cfg.prevLocationAreaUrl = resp.Previous
	cfg.nextLocationAreaUrl = resp.Next

	return nil
}

func callbackMapb(cfg *config, args ...string) error {
	if cfg.prevLocationAreaUrl == nil {
		return errors.New("No previous location area URL")
	}

	c := cfg.pokeapiClient

	resp, err := c.ListLocationAreas(cfg.prevLocationAreaUrl)

	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, v := range resp.Results {
		fmt.Printf(" - %s\n", v.Name)
	}

	cfg.prevLocationAreaUrl = resp.Previous
	cfg.nextLocationAreaUrl = resp.Next

	return nil
}
