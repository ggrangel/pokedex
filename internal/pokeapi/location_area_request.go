package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResponse, error) {
	var fullUrl string
	if pageUrl != nil {
		fullUrl = *pageUrl
	} else {
		endpoint := "/location-area"
		fullUrl = baseUrl + endpoint
	}

	dat, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Println("Cache hit")
		locationAreasResponse := LocationAreasResponse{}
		err := json.Unmarshal(dat, &locationAreasResponse)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return locationAreasResponse, nil
	}
	fmt.Println("Cache miss")

	req, err := http.NewRequest("GET", fullUrl, nil)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	if resp.StatusCode >= 400 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreasResponse := LocationAreasResponse{}
	err = json.Unmarshal(dat, &locationAreasResponse)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(fullUrl, dat)

	return locationAreasResponse, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullUrl := baseUrl + endpoint

	dat, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Println("Cache hit")
		locationAreasResponse := LocationArea{}
		err := json.Unmarshal(dat, &locationAreasResponse)
		if err != nil {
			return LocationArea{}, err
		}
		return locationAreasResponse, nil
	}
	fmt.Println("Cache miss")

	req, err := http.NewRequest("GET", fullUrl, nil)

	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationArea{}, err
	}

	if resp.StatusCode >= 400 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return LocationArea{}, err
	}

	locationAreasResponse := LocationArea{}
	err = json.Unmarshal(dat, &locationAreasResponse)

	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullUrl, dat)

	return locationAreasResponse, nil
}
