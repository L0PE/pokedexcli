package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + name
	
	data, ok := c.cache.Get(url)
	if !ok {
		response, err := c.httpClient.Get(url)
		if err != nil {
			return Pokemon{}, fmt.Errorf("Error during making reuest: %w", err)
		}

		defer response.Body.Close()

		data, err = io.ReadAll(response.Body)
		if err != nil {
			return Pokemon{}, fmt.Errorf("Error during reading data: %w", err)
		}

		c.cache.Add(url, data)
 	}

	var pokeapiResponse Pokemon
	if err := json.Unmarshal(data, &pokeapiResponse); err != nil {
		return Pokemon{}, fmt.Errorf("Error during unmarshaling data: %w", err)
	}

	return pokeapiResponse, nil 
}

