package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ForecastData struct {
	Time          []string  `json:"time"`
	Temperature2m []float64 `json:"temperature_2m"`
}

type WeatherResponse struct {
	TimezoneAbbrev string       `json:"timezone_abbreviation"`
	Hourly         ForecastData `json:"hourly"`
}

func FetchWeather(latitude float64, longitude float64) WeatherResponse {
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&hourly=temperature_2m", latitude, longitude)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("Could not fetch data")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var weatherResponse WeatherResponse
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		log.Fatal(err)
	}

	return weatherResponse
}
