package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationDetails(location string) (LocationData, error) {
	url := baseURL + "/location-area/" + location

	if data, exists := c.cache.Get(url); exists {
		locationResp := LocationData{}
		err := json.Unmarshal(data, &locationResp)
		if err != nil {
			return LocationData{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationData{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationData{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationData{}, err
	}

	detailsResp := LocationData{}
	err = json.Unmarshal(dat, &detailsResp)
	if err != nil {
		return LocationData{}, err
	}

	c.cache.Add(url, dat)
	return detailsResp, nil
}
