package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
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

func GetLocation() UserLocation {
	loc := strings.Join(os.Args[1:], " ")

	r, err := http.Get("https://geocoding-api.open-meteo.com/v1/search?name=" + url.QueryEscape(loc))
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

	location := unmarshaledResponse.Results[0]

	return UserLocation{location.Name, location.Latitude, location.Longitude, location.Timezone}
}

type UserLocation struct {
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string
}

func GetUserLocation() UserLocation {
	ipapiClient := http.Client{}
	req, err := http.NewRequest("GET", "https://ipapi.co/json/", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.")
	resp, err := ipapiClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var userLocation UserLocation
	json.Unmarshal(body, &userLocation)

	return userLocation
}

func (l UserLocation) PrintLoadingMessage() {
	msg := fmt.Sprint("Fetching data for ", l.City, "...")
	color.Green(msg)
}
