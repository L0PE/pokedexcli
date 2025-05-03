package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ExploreLocation(loation_name string) (ExploreLocationReponse, error) {
	url := baseUrl + "/location-area/" + loation_name
	
	data, ok := c.cache.Get(url)
	if !ok {
		response, err := c.httpClient.Get(url)
		if err != nil {
			return ExploreLocationReponse{}, fmt.Errorf("Error during making reuest: %w", err)
		}

		defer response.Body.Close()

		data, err = io.ReadAll(response.Body)
		if err != nil {
			return ExploreLocationReponse{}, fmt.Errorf("Error during reading data: %w", err)
		}

		c.cache.Add(url, data)
 	}

	var pokeapiResponse ExploreLocationReponse
	if err := json.Unmarshal(data, &pokeapiResponse); err != nil {
		return ExploreLocationReponse{}, fmt.Errorf("Error during unmarshaling data: %w", err)
	}

	return pokeapiResponse, nil 
}

