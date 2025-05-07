package main

import "fmt"

func commandPokedex(conf *config, _ []string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range conf.Pokedex {
		fmt.Printf("\t-%s\n", pokemon.Name)
	}

	return nil
}
