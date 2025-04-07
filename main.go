package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println("Hello, World!")
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
