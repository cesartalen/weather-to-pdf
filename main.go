package main

import (
	"log"
	"os"

	"github.com/cesartalen/weather-to-pdf/pdf"
	"github.com/cesartalen/weather-to-pdf/weather"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Use param <city>")
	}
	city := os.Args[1]

	var locationData weather.Location
	var weatherData weather.WeatherResponse

	locationData = weather.FetchLocation(city)
	weatherData = weather.FetchWeather(locationData.Latitude, locationData.Longitude)

	err := pdf.GeneratePDF("weather.pdf", weatherData, locationData)
	if err != nil {
		log.Fatal(err)
	}
}
