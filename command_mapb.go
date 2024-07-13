package main

import (
	"errors"
	"fmt"

	"github.com/scottEAdams1/REPLPokedex/internal/pokeapi"
)

func commandMapB(cfg *config, extra string) error {
	if cfg.previous == nil {
		return errors.New("you're on the first page")
	}
	locations, err := pokeapi.ListLocations(*cfg.previous, cfg.cache)
	if err != nil {
		return err
	}
	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}
	cfg.next = locations.Next
	cfg.previous = locations.Previous
	return nil
}
