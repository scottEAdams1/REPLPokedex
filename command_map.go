package main

import (
	"fmt"

	"github.com/scottEAdams1/REPLPokedex/internal/pokeapi"
)

func commandMap(cfg *config, extra string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.next != nil {
		url = *cfg.next
	}
	locations, err := pokeapi.ListLocations(url, cfg.cache)
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
