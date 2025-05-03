package main

import (
	"fmt"

	"github.com/L0PE/pokedexcli/internal/pokeapi"
)

func commandMap(conf *config, _ []string) error  {
	return baseMap(conf, conf.Next)
}

func commandMapb(conf *config, _ []string) error {
	if conf.Previous == nil {
		return fmt.Errorf("You are allready on a friest page")
	}

	return baseMap(conf, conf.Previous)
}

func baseMap(conf *config, url *string) error {
	poreResponse, err := conf.PokeapiClient.ListLocation(url)
	if err != nil {
		return err
	}

	conf.Next = poreResponse.Next
	conf.Previous = poreResponse.Previous
	printNames(poreResponse.Results)

	return nil
}

func printNames(results []pokeapi.PokeapiResult) {
	for _, result := range results {
		fmt.Println(result.Name)
	}
}
