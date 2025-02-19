package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/KOTBCAnorax/pokedex/internal/pokecache"
)

type PokeArea struct {
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

var Config = config{Prev: "", Next: "https://pokeapi.co/api/v2/location-area/"}

func makePokeHttpRequest(url string, c *pokecache.Cache) error {
	body, ok := c.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("PokeAPI http request failed: %v", err)
		}

		body, err = io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return fmt.Errorf("failed to read PokeAPI http response: %v", err)
		}

		c.Add(url, body)
	}

	locations := PokeArea{}
	err := json.Unmarshal(body, &locations)
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
	makePokeHttpRequest(Config.Next, c)
	return nil
}

func RetreatMap(c *pokecache.Cache) error {
	if Config.Prev == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	makePokeHttpRequest(Config.Prev, c)
	return nil
}
