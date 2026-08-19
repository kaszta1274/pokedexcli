package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/kaszta1274/pokedexcli/internal/pokecache"
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
	cache   *pokecache.Cache
}

func NewClient() *Client {
	return &Client{
		baseURL: "https://pokeapi.co/api/v2/location-area/",
		cache:   pokecache.NewCache(time.Hour),
	}
}

func (c *Client) GetMap(url *string) (LocationAreas, error) {
	reqURL := c.baseURL
	if url != nil {
		reqURL = *url
	}

	if cachedBytes, exists := c.cache.Get(reqURL); exists {
		locationAreas, err := parseLocationAreas(cachedBytes)
		if err != nil {
			return LocationAreas{}, err
		}

		return locationAreas, nil
	}

	responseBody, err := c.getBody(reqURL)
	if err != nil {
		return LocationAreas{}, err
	}

	c.cache.Add(reqURL, responseBody)

	return parseLocationAreas(responseBody)
}

func (c *Client) getBody(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code %d: %s\n", res.StatusCode, responseBody)
	}

	return responseBody, nil
}

func parseLocationAreas(cachedBytes []byte) (LocationAreas, error) {
	var locationAreas LocationAreas
	if err := json.Unmarshal(cachedBytes, &locationAreas); err != nil {
		return LocationAreas{}, err
	}

	return locationAreas, nil
}
