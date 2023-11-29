package main

import "github.com/grepexdev/go-pokedex/internal/pokeapi"

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config {
		pokeapiClient: pokeapi.NewClient(),
	}
	startREPL(&cfg)
}
