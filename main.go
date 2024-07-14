package main

import (
	"time"

	"github.com/scottEAdams1/REPLPokedex/internal/pokeapi"
	"github.com/scottEAdams1/REPLPokedex/internal/pokecache"
)

func main() {
	cfg := &config{cache: pokecache.NewCache(5 * time.Minute), pokedex: make(map[string]pokeapi.Pokemon)}
	startRepl(cfg)
}
