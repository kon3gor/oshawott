package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	url      = "https://pokeapi.co/api/v2/pokemon?limit=100000&offset=0"
	keysFile = "keys.txt"
)

type response struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func main() {
	keys := fetchKeys()
	fmt.Printf("Fetcehd %d keys", len(keys))
	err := os.WriteFile(keysFile, []byte(strings.Join(keys, ",")), 644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func fetchKeys() []string {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer res.Body.Close()

	var pr response
	if err := json.NewDecoder(res.Body).Decode(&pr); err != nil {
		return make([]string, 0)
	}

	keys := make([]string, len(pr.Results))
	for i, p := range pr.Results {
		keys[i] = p.Name
	}

	return keys
}
