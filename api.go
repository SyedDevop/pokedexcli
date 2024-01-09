package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Pokedex struct {
	Previous *string  `json:"previous,omitempty"`
	Next     string   `json:"next"`
	Results  []Result `json:"results"`
	Count    int64    `json:"count"`
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// func (p *Pokedex) String() []string {
// 	var data []string
// 	for _, result := range p.Results {
// 		data = append(data, result.Name)
// 	}
// 	return data
// }

type Clint struct {
	HTTPClint *http.Client
	prevUri   *string
	cacheData map[string]Pokedex
	nextUri   string
}

func NewClient() *Clint {
	return &Clint{
		nextUri:   "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		prevUri:   nil,
		HTTPClint: &http.Client{},
		cacheData: make(map[string]Pokedex),
	}
}

func (c *Clint) sendRequest(req *http.Request, v interface{}) error {
	res, err := c.HTTPClint.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, v); err != nil {
		return err
	}
	return nil
}

// FIX : bug wen I'm going back to first page
// and then I try to get data from the next page
// i don't get the initial page 0
func (c *Clint) GetPokeList() (*Pokedex, error) {
	cacheData, dataExists := c.cacheData[c.nextUri]
	if dataExists {
		c.nextUri = cacheData.Next
		c.prevUri = cacheData.Previous
		return &cacheData, nil
	}
	req, err := http.NewRequest(http.MethodGet, c.nextUri, nil)
	if err != nil {
		return nil, err
	}
	var data Pokedex

	err = c.sendRequest(req, &data)
	if err != nil {
		return nil, err
	}

	c.cacheData[c.nextUri] = data
	c.nextUri = data.Next
	c.prevUri = data.Previous
	return &data, nil
}

func (c *Clint) GetPokePrevesList() (*Pokedex, error) {
	cacheData, dataExists := c.cacheData[*c.prevUri]
	if dataExists {
		c.nextUri = cacheData.Next
		c.prevUri = cacheData.Previous
		return &cacheData, nil
	}
	req, err := http.NewRequest(http.MethodGet, *c.prevUri, nil)
	if err != nil {
		return nil, err
	}
	var data Pokedex

	err = c.sendRequest(req, &data)
	if err != nil {
		return nil, err
	}
	c.nextUri = data.Next
	c.prevUri = data.Previous
	return &data, nil
}
