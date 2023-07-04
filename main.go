package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/fatih/color"
)

type Weather struct {
	Temperature string `json:"temperature"`
	Wind        string `json:"wind"`
	Description string `json:"description"`
	Forecast    []struct {
		Day         string `json:"day"`
		Temperature string `json:"temperature"`
		Wind        string `json:"wind"`
	} `json:"forecast"`
}

func main() {
	location := "Mairipora"

	if len(os.Args) >= 2 {
		location = os.Args[1]
	}

	res, err := http.Get("https://goweather.herokuapp.com/weather/" + location)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	json.Unmarshal(body, &weather)

	temperature, wind, description, forecast := weather.Temperature, weather.Wind, weather.Description, weather.Forecast

	title := fmt.Sprintf("%s: %s\n%s, Vento: %s\n", location, temperature, description, wind)
	weather.printWithColor(title)

	for _, day := range forecast {
		fmt.Printf("%s - %s - Vento: %s\n", day.Day, day.Temperature, day.Wind)
	}
}

func (w Weather) printWithColor(s string) {
	tempString := w.Temperature[1:3]

	temperature, err := strconv.Atoi(tempString)
	if err != nil {
		panic(err)
	}

	if temperature >= 30 {
		color.Red(s)
		return
	} else if temperature < 30 && temperature > 15 {
		color.Yellow(s)
	} else {
		color.Blue(s)
	}

}
