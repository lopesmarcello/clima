package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
)

type Weather struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Hourly    struct {
		Time        []int64   `json:"time"`
		Temperature []float32 `json:"temperature_2m"`
	} `json:"hourly"`
	Daily struct {
		Time           []int64   `json:"time"`
		TemperatureMin []float32 `json:"temperature_2m_min"`
		TemperatureMax []float32 `json:"temperature_2m_max"`
		Sunrise        []int64   `json:"sunrise"`
		Sunset         []int64   `json:"sunset"`
	}
}

func (l Location) GetWeather() Weather {
	latitude, longitude := fmt.Sprintf("%f", l.Latitude), fmt.Sprintf("%f", l.Longitude)

	// r, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=" + latitude + "&longitude=" + longitude + "&hourly=temperature_2m,rain&daily=weathercode,temperature_2m_max,temperature_2m_min,apparent_temperature_max,apparent_temperature_min,sunrise,sunset,uv_index_max,uv_index_clear_sky_max,precipitation_sum,rain_sum,showers_sum,snowfall_sum,precipitation_hours,precipitation_probability_max,windspeed_10m_max,windgusts_10m_max,winddirection_10m_dominant,shortwave_radiation_sum,et0_fao_evapotranspiration&timezone=America/Sao_Paulo")
	r, err := http.Get("https://api.open-meteo.com/v1/dwd-icon?latitude=" + latitude + "&longitude=" + longitude + "&hourly=temperature_2m&daily=temperature_2m_max,temperature_2m_min,sunrise,sunset&timeformat=unixtime&timezone=America%2FSao_Paulo&format=json")
	if err != nil {
		log.Fatal("Error getting weather", err)
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		log.Fatal("Weather API not available")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var unmarshaledResponse Weather
	json.Unmarshal(body, &unmarshaledResponse)

	return unmarshaledResponse
}

func (w Weather) PrintHourlyWeather() {
	t, temperature := w.Hourly.Time, w.Hourly.Temperature
	day := time.Unix(t[0], 0).Format("02/01/2006")

	dayMsg := fmt.Sprintf("Day: %s", day)
	color.Yellow(dayMsg)

	var hourTempMsg string
	var currentTemp float32

	for i, hour := range t {
		date := time.Unix(hour, 0)

		if date.Hour() == time.Now().Hour() && date.Day() == time.Now().Day() {
			currentTemp = temperature[i]
			hourTempMsg = fmt.Sprintf("%sh - %.1f°C\n\n", date.Format("15:04"), temperature[i])
		}
	}

	if currentTemp >= 27 {
		color.Red(hourTempMsg)
	} else if currentTemp < 27 && currentTemp > 17 {
		color.Yellow(hourTempMsg)
	} else {
		color.Blue(hourTempMsg)
	}
}

func (w Weather) PrintDailyWeather() {
	header := fmt.Sprintln("Day  ", "-", "Min  ", "-", "Max  ", "-", "Sunrise", "-", "Sunset")
	color.Blue(header)
	for i, day := range w.Daily.Time {
		date := time.Unix(day, 0)
		formattedDay := date.Format("02/01")
		minTemp := fmt.Sprintf("%.1f°C", w.Daily.TemperatureMin[i])
		maxTemp := fmt.Sprintf("%.1f°C", w.Daily.TemperatureMax[i])
		sunrise := time.Unix(w.Daily.Sunrise[i], 0).Format("15:04")
		sunset := time.Unix(w.Daily.Sunset[i], 0).Format("15:04")

		fmt.Println(formattedDay, "-", minTemp, "-", maxTemp, "-", sunrise, "-", sunset)
	}
}
