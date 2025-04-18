package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/L0PE/pokedexcli/internal/pokeapi"
)

type cliCommands struct {
	name        string
	description string
	callback    func(conf *config) error
}

var availableCommands map[string]cliCommands

func init() {
	availableCommands = map[string]cliCommands{
		"exit": {
			name:        "exit",
			description: "Exit the pockedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name: "map",
			description: "Get names of local Areas",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Get names of local Areas",
			callback: commandMapb,
		},
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	conf := config{
		PokeapiClient: pokeapi.NewClient(5 * time.Second),
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		result := scanner.Text()
		clearedInput := cleanInput(result)
		firstWord := ""
		if len(clearedInput) >= 1 {
			firstWord = clearedInput[0]
		}

		command, ok := availableCommands[firstWord]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(&conf)
		if err != nil {
			fmt.Printf("Error executing command: %v \n", err)
		}
	}
}

func commandExit(_ *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, v := range availableCommands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}

func cleanInput(text string) []string {
	text = strings.Trim(text, " ")
	textSlice := strings.Split(text, " ")
	var returnSlice []string

	for _, v := range textSlice {
		v = strings.Trim(v, " ")
		if v == "" {
			continue
		}

		returnSlice = append(returnSlice, strings.ToLower(v))
	}

	return returnSlice
}
