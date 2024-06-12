package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullUrl := baseUrl + endpoint

	dat, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Println("Cache hit")
		pokemon := Pokemon{}
		err := json.Unmarshal(dat, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}
	fmt.Println("Cache miss")

	req, err := http.NewRequest("GET", fullUrl, nil)

	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return Pokemon{}, err
	}

	if resp.StatusCode >= 400 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(dat, &pokemon)

	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullUrl, dat)

	return pokemon, nil
}
