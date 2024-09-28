package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint

	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	//check cache
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		//cache hit
		var locationAreas LocationAreaResp
		if err := json.Unmarshal(dat, &locationAreas); err != nil {
			return LocationAreaResp{}, err
		}
	}

	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)

	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreaResp{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return LocationAreaResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return LocationAreaResp{}, err
	}

	var locationAreas LocationAreaResp
	if err := json.Unmarshal(data, &locationAreas); err != nil {
		return LocationAreaResp{}, err
	}

	c.cache.Add(fullUrl, data)

	return locationAreas, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullUrl := baseUrl + endpoint

	//check cache
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		//cache hit
		var locationAreas LocationAreaResp
		if err := json.Unmarshal(dat, &locationAreas); err != nil {
			return LocationArea{}, err
		}
	}

	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)

	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return LocationArea{}, err
	}

	var locationArea LocationArea
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullUrl, data)

	return locationArea, nil
}
