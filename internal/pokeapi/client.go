package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Client struct {
	baseURL string
}

func NewClient() *Client {
	return &Client{
		baseURL: "https://pokeapi.co/api/v2/location-area/",
	}
}

func (c *Client) GetMap(url *string) (LocationAreas, error) {
	reqURL := c.baseURL
	if url != nil {
		reqURL = *url
	}

	res, err := http.Get(reqURL)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	if res.StatusCode > 299 {
		return LocationAreas{}, fmt.Errorf("response failed with status code %d: %s\n", res.StatusCode, responseBody)
	}

	var locationAreas LocationAreas
	if err = json.Unmarshal(responseBody, &locationAreas); err != nil {
		return LocationAreas{}, err
	}

	return locationAreas, nil
}
