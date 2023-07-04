package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
)

type Location struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Country   string  `json:"country"`
}

type Response struct {
	Results []Location `json:"results"`
}

func GetLocation() Location {
	loc := "mairipora"

	if len(os.Args) > 1 {
		loc = strings.Join(os.Args[1:], "+")
	}

	r, err := http.Get("https://geocoding-api.open-meteo.com/v1/search?name=" + loc)
	if err != nil {
		log.Fatal("Error retrieving location", err)
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		log.Fatal("Location API not available")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var unmarshaledResponse Response
	json.Unmarshal(body, &unmarshaledResponse)

	return unmarshaledResponse.Results[0]
}

func (l Location) PrintLoadingMessage() {
	msg := fmt.Sprint("Fetching data for ", l.Name, "...")
	color.Green(msg)
}
