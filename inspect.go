package main

import (
	"fmt"
)

func commandInspect(conf *config, arguments []string) error {
	if len(arguments) < 1 {
		return fmt.Errorf("No pokemon name provided. Please provide one argument.")
	}	

	name := arguments[0]
	pokemon, ok := conf.Pokedex[name]

	if !ok {
		return fmt.Errorf("Pokemon is not in the Pokedex")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Printf("Stats:\n")
	for _, pokemon_stat := range pokemon.Stats {
		fmt.Printf("\t-%s: %d\n", pokemon_stat.Stat.Name, pokemon_stat.Value)
	}

	fmt.Printf("Types:\n")
	for _, pokemon_type := range pokemon.Types {
		fmt.Printf("\t-%s\n", pokemon_type.Type.Name)
	}

	return nil
}
