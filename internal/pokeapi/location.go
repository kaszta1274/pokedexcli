package pokeapi

import "encoding/json"

type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func parseRespShallowLocations(cachedBytes []byte) (RespShallowLocations, error) {
	var locationsResp RespShallowLocations
	if err := json.Unmarshal(cachedBytes, &locationsResp); err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}

type RespDetailedLocation struct {
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func parseRespDetailedLocation(cachedBytes []byte) (RespDetailedLocation, error) {
	var locationResp RespDetailedLocation
	if err := json.Unmarshal(cachedBytes, &locationResp); err != nil {
		return RespDetailedLocation{}, err
	}

	return locationResp, nil
}
