package pokeAPI

import (
	"encoding/json"
	"fmt"

	"github.com/KOTBCAnorax/pokedex/internal/pokecache"
)

type PokeArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetPokemonsList(areaname string, c *pokecache.Cache) error {
	url := LocationAreaURL + "/" + areaname
	response, err := makePokeHttpRequest(url, c)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	pokemons := PokeArea{}
	err = json.Unmarshal(response, &pokemons)
	if err != nil {
		return fmt.Errorf("failed to decode http response: %v", err)
	}

	for i := range pokemons.PokemonEncounters {
		fmt.Println(pokemons.PokemonEncounters[i].Pokemon.Name)
	}

	return nil
}
