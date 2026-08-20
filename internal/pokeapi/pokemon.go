package pokeapi

import "encoding/json"

type Pokemon struct {
	Name           string `json:"name"`
	ID             int    `json:"id"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func parsePokemon(bytes []byte) (Pokemon, error) {
	var pokemon Pokemon
	if err := json.Unmarshal(bytes, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
