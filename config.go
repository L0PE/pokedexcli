package main

import "github.com/L0PE/pokedexcli/internal/pokeapi"

type config struct {
	PokeapiClient pokeapi.Client
	Next 			*string
	Previous 	*string
	Pokedex 	map[string]pokeapi.Pokemon
}
