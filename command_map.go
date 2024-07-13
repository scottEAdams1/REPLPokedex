package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(cfg *config) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if cfg.next != nil {
		res, err = http.Get(*cfg.next)
	}
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return err
	}
	if err != nil {
		return err
	}
	locations := Locations{}
	err1 := json.Unmarshal(body, &locations)
	if err1 != nil {
		return err1
	}
	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}
	cfg.next = locations.Next
	cfg.previous = locations.Previous
	return nil
}

type Locations struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}
type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
