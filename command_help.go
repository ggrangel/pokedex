package main

import (
	"fmt"
)

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Here are your available commands: ")
	for _, v := range getCommands() {
		fmt.Println(v.name, "-", v.description)
	}

	return nil
}
