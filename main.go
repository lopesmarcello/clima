package main

import "fmt"

func main() {
	location := GetLocation()
	location.PrintLoadingMessage()

	fmt.Println("")

	weather := location.GetWeather()

	weather.PrintHourlyWeather()
	weather.PrintDailyWeather()
}
