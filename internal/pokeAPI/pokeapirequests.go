package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func makePokeHttpRequest(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("PokeAPI http request failed: %v", err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return fmt.Errorf("failed to read PokeAPI http response: %v", err)
	}

	locations := PokeArea{}
	err = json.Unmarshal(body, &locations)
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

func AdvanceMap() error {
	if Config.Next == "" {
		fmt.Println("you're on the last page")
		return nil
	}
	makePokeHttpRequest(Config.Next)
	return nil
}

func RetreatMap() error {
	if Config.Prev == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	makePokeHttpRequest(Config.Prev)
	return nil
}
