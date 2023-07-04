package main

import (
	"fmt"
	"os"
)

func main() {
	var location UserLocation

	if len(os.Args) > 1 {
		location = GetLocation()
	} else {
		location = GetUserLocation()
	}

	location.PrintLoadingMessage()

	fmt.Printf("Timezone: America/Sao Paulo\n\n")

	weather := location.GetWeather()
	weather.PrintHourlyWeather()
	weather.PrintDailyWeather()
}
