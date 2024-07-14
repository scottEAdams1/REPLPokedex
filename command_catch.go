package main

import (
	"fmt"
	"math/rand"

	"github.com/scottEAdams1/REPLPokedex/internal/pokeapi"
)

func commandCatch(cfg *config, name string) error {
	pokemon, err := pokeapi.PokemonStats("https://pokeapi.co/api/v2/pokemon/"+name+"/", cfg.cache)
	if err != nil {
		return err
	}
	res := rand.Intn(pokemon.BaseExperience)
	if res > 40 {
		fmt.Println(name + " escaped!")
		return nil
	}
	fmt.Println(name + " was caught!")
	cfg.pokedex[pokemon.Name] = pokemon
	return nil

}
