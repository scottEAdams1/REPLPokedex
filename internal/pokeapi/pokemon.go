package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/scottEAdams1/REPLPokedex/internal/pokecache"
)

func PokemonStats(url string, c pokecache.Cache) (Pokemon, error) {
	pokemon := Pokemon{}
	body, ok := c.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return pokemon, err
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return pokemon, err
		}
		if err != nil {
			return pokemon, err
		}
	}

	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		return pokemon, err
	}
	c.Add(url, body)
	return pokemon, nil
}
