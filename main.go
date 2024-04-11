package main

import (
	"cesartalen/weather-to-pdf/pdf"
	"cesartalen/weather-to-pdf/weather"
	"log"
)

func main() {
	var locationData weather.Location
	var weatherData weather.WeatherResponse

	locationData = weather.FetchLocation("Visby")
	weatherData = weather.FetchWeather(locationData.Latitude, locationData.Longitude)

	err := pdf.GeneratePDF("weather.pdf", weatherData)
	if err != nil {
		log.Fatal(err)
	}
}
