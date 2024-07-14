package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/scottEAdams1/REPLPokedex/internal/pokeapi"
	"github.com/scottEAdams1/REPLPokedex/internal/pokecache"
)

type config struct {
	next     *string
	previous *string
	cache    pokecache.Cache
	pokedex  map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()
		text = strings.ToLower(text)
		words := strings.Fields(text)
		commandName := strings.ToLower(words[0])
		commandMap := getCommands()
		item, ok := commandMap[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		var err error
		if len(words) == 1 {
			err = item.callback(cfg, "")
		} else {
			err = item.callback(cfg, words[1])
		}
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore <area_name>",
			description: "Displays the names of the pokemon at that location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Throws a pokeball to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Shows stats of that pokemon if you have caught it",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays the names of all caught pokemon",
			callback:    commandPokedex,
		},
	}
}
