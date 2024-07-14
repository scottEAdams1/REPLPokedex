package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, name string) error {
	pokemon, ok := cfg.pokedex[name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, type1 := range pokemon.Types {
		fmt.Println("  - " + type1.Type.Name)
	}
	return nil
}
