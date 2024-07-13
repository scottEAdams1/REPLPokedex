package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandMapB(cfg *config) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if cfg.previous == nil {
		return errors.New("you're on the first page")
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
