package main

import "fmt"

func commandExplore(conf *config, argumets []string) error {
		if len(argumets) < 1 {
			return fmt.Errorf("No area provided. Please provide one argument.")
		}	

		name := argumets[0]

		response, err := conf.PokeapiClient.ExploreLocation(name) //TODO: update with actual
		if err != nil {
			return err
		}

		fmt.Printf("Exploring %s\n", name)
		fmt.Printf("Found Pokemon:\n")

		for _, encounter := range response.PokemonEncounters {
			fmt.Printf("\t- %s\n", encounter.Pokemon.Name)
		}

		return nil
}
