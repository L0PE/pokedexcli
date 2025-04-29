package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListLocation(pageURL *string) (PokeapiReponse, error) {
	url := baseUrl + "/location-area"


	if pageURL != nil {
		url = *pageURL
	}
	
	data, ok := c.cache.Get(url)
	if !ok {
		response, err := c.httpClient.Get(url)
		if err != nil {
			return PokeapiReponse{}, fmt.Errorf("Error during making reuest: %w", err)
		}

		defer response.Body.Close()

		data, err = io.ReadAll(response.Body)
		if err != nil {
			return PokeapiReponse{}, fmt.Errorf("Error during reading data: %w", err)
		}

		c.cache.Add(url, data)
 	}

	var pokeapiResponse PokeapiReponse
	if err := json.Unmarshal(data, &pokeapiResponse); err != nil {
		return PokeapiReponse{}, fmt.Errorf("Error during unmarshaling data: %w", err)
	}

	return pokeapiResponse, nil 
}

