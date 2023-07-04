package main

import "fmt"

func main() {
	location := GetLocation()
	location.PrintLoadingMessage()

	fmt.Printf("Timezone: America/Sao Paulo\n\n")

	weather := location.GetWeather()
	weather.PrintHourlyWeather()
	weather.PrintDailyWeather()
}
