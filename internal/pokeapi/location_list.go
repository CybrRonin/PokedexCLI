package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if data, exists := c.cache.Get(url); exists {
		locationsResp := Locations{}
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return Locations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Locations{}, err
	}

	locationsResp := Locations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return Locations{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
