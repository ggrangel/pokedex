package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" >")

		scanner.Scan()

		text := scanner.Text()

		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		callback, ok := getCommands()[commandName]

		if !ok {
			fmt.Println("Command not found")
			continue
		}

		err := callback.callback(cfg, args...)

		if err != nil {
			fmt.Println("Error running command: ", err)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "displays this help menu",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "displays the map",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the previous map",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "explore a location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon}",
			description: "catch a pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon}",
			description: "print information about a pokemon",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "print the pokedex",
			callback:    callbackPokedex,
		},
		"exit": {
			name:        "exit",
			description: "exits the program",
			callback:    callbackExit,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
