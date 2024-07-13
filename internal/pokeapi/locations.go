package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func ListLocations(url string) (Locations, error) {
	locations := Locations{}
	res, err := http.Get(url)
	if err != nil {
		return locations, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return locations, err
	}
	if err != nil {
		return locations, err
	}

	err1 := json.Unmarshal(body, &locations)
	if err1 != nil {
		return locations, err1
	}
	return locations, nil
}
