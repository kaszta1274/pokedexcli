package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/kaszta1274/pokedexcli/internal/pokecache"
)

type Client struct {
	baseURL string
	cache   *pokecache.Cache
}

func NewClient() *Client {
	return &Client{
		baseURL: "https://pokeapi.co/api/v2/",
		cache:   pokecache.NewCache(5 * time.Minute),
	}
}

func (c *Client) GetLocations(pageURL *string) (RespShallowLocations, error) {
	reqURL := c.baseURL + "location-area/"
	if pageURL != nil {
		reqURL = *pageURL
	}

	if cachedBytes, exists := c.cache.Get(reqURL); exists {
		locationsResp, err := parseRespShallowLocations(cachedBytes)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil
	}

	responseBody, err := c.getBody(reqURL)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(reqURL, responseBody)

	return parseRespShallowLocations(responseBody)
}

func (c *Client) GetLocationDetails(locationName string) (RespDetailedLocation, error) {
	reqURL := c.baseURL + "location-area/" + locationName

	if cachedBytes, exists := c.cache.Get(reqURL); exists {
		detailedLocationResponse, err := parseRespDetailedLocation(cachedBytes)
		if err != nil {
			return RespDetailedLocation{}, err
		}

		return detailedLocationResponse, nil
	}

	responseBody, err := c.getBody(reqURL)
	if err != nil {
		return RespDetailedLocation{}, err
	}

	c.cache.Add(reqURL, responseBody)

	return parseRespDetailedLocation(responseBody)
}

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	reqURL := c.baseURL + "pokemon/" + pokemonName

	responseBody, err := c.getBody(reqURL)
	if err != nil {
		return Pokemon{}, err
	}

	return parsePokemon(responseBody)
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
