package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Location struct {
	Name        string  `json:"name"`
	Country     string  `json:"country"`
	CountryCode string  `json:"country_code"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
}

type LocationResponse struct {
	Results []Location `json:"results"`
}

func FetchLocation(city string) Location {
	url := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1", city)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Could not fetch data")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var locationResp LocationResponse
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		log.Fatal(err)
	}

	// TODO handle this in a better way?
	location := Location{
		Name:        locationResp.Results[0].Name,
		Country:     locationResp.Results[0].Country,
		CountryCode: locationResp.Results[0].CountryCode,
		Longitude:   locationResp.Results[0].Longitude,
		Latitude:    locationResp.Results[0].Latitude,
	}

	return location
}
