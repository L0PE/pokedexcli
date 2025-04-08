package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		result := scanner.Text()
		clearedInput := cleanInput(result)
		firstWord := ""
		if len(clearedInput) >= 1 {
			firstWord = clearedInput[0]
		}	
		fmt.Println("Your command was: " + firstWord)
	}
}

func cleanInput(text string) []string  {
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
