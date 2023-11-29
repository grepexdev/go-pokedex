package main

import (
	"fmt"
	"log"

	"github.com/grepexdev/go-pokedex/internal/pokeapi"
)

func callbackMap() error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreasResp()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location areas: ")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	return nil
}
