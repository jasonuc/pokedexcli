package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(locationAreaName string) (Pokemon, error) {
	endpoint := "/pokemon/" + locationAreaName
	fullUrl := baseUrl + endpoint

	//check cache
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		//cache hit
		var pokemon Pokemon
		if err := json.Unmarshal(dat, &pokemon); err != nil {
			return Pokemon{}, err
		}
	}

	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)

	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullUrl, data)

	return pokemon, nil
}
