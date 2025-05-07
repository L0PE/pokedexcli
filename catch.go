package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(conf *config, arguments []string) error {
	if len(arguments) < 1 {
		return fmt.Errorf("No pokemon name provided. Please provide one argument.")
	}	

	name := arguments[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := conf.PokeapiClient.GetPokemon(name)

	if err != nil {
		return fmt.Errorf("Error while getting pokemon details %w", err)
	}

	rand := rand.Int63n(pokemon.BaseExperience)

	if rand > pokemon.BaseExperience / 3 {
		fmt.Printf("%s was caught!\n", name)
		fmt.Println("You may now inspect it with the inspect command.")
		conf.Pokedex[name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
