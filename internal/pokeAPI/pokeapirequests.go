package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/KOTBCAnorax/pokedex/internal/pokecache"
)

type PokeLocations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type config struct {
	Prev string
	Next string
}

var LocationAreaURL = "https://pokeapi.co/api/v2/location-area/"
var Config = config{Prev: "", Next: LocationAreaURL}

func makePokeHttpRequest(url string, c *pokecache.Cache) ([]byte, error) {
	body, ok := c.Get(url)
	if !ok {
		res, err := http.Get(url)
		if res.StatusCode == 404 {
			return nil, fmt.Errorf("PokeAPI http request failed: 404 not found, please verify your input")
		}

		if err != nil {
			return nil, fmt.Errorf("PokeAPI http request failed: %v", err)
		}

		body, err = io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("failed to read PokeAPI http response: %v", err)
		}

		c.Add(url, body)
	}

	return body, nil
}

func GetLocationsList(response []byte) error {
	locations := PokeLocations{}
	err := json.Unmarshal(response, &locations)
	if err != nil {
		return fmt.Errorf("failed to decode http response: %v", err)
	}

	Config.Prev = locations.Previous
	Config.Next = locations.Next

	for i := range locations.Results {
		fmt.Println(locations.Results[i].Name)
	}

	return nil
}

func AdvanceMap(c *pokecache.Cache) error {
	if Config.Next == "" {
		fmt.Println("you're on the last page")
		return nil
	}
	response, _ := makePokeHttpRequest(Config.Next, c)
	GetLocationsList(response)
	return nil
}

func RetreatMap(c *pokecache.Cache) error {
	if Config.Prev == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	response, _ := makePokeHttpRequest(Config.Prev, c)
	GetLocationsList(response)
	return nil
}
