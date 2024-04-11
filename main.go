package main

import (
	"cesartalen/weather-to-pdf/pdf"
	"log"
)

func main() {
	err := pdf.GeneratePDF("weather.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
