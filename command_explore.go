package main

import (
	"fmt"

	"github.com/scottEAdams1/REPLPokedex/internal/pokeapi"
)

func commandExplore(cfg *config, name string) error {
	pokemon, err := pokeapi.ListPokemon("https://pokeapi.co/api/v2/location-area/"+name+"/", cfg.cache)
	if err != nil {
		return err
	}
	for _, result := range pokemon.PokemonEncounters {
		fmt.Println(" - ", result.Pokemon.Name)
	}
	return nil
}
