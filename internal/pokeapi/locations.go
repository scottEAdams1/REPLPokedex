package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/scottEAdams1/REPLPokedex/internal/pokecache"
)

func ListLocations(url string, c pokecache.Cache) (Locations, error) {
	locations := Locations{}
	body, ok := c.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return locations, err
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return locations, err
		}
		if err != nil {
			return locations, err
		}
	}
	err := json.Unmarshal(body, &locations)
	if err != nil {
		return locations, err
	}
	c.Add(url, body)
	return locations, nil
}
