package pdf

import (
	"cesartalen/weather-to-pdf/weather"
	"fmt"
	"log"
	"time"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/core"
)

func GeneratePDF(filename string, weatherData weather.WeatherResponse) error {
	m := maroto.New()

	// For every day, divides by 24 to get count of days.
	for i := 0; i < (len(weatherData.Hourly.Time) / 24); i++ {
		tempCols := make([]core.Col, 0, 7)
		timeCols := make([]core.Col, 0, 7)

		/* j*3 because 24/8 = 3, dividing day into 8 parts. Loops 8 times for every days worth of data. Should not need to be modified if more days are added.
		Doing this to make the data fit better on the page.
		Stores data in two cols so time can be put in a row above temp.
		*/
		for j := 0; j <= 7; j++ {
			index := i*24 + j*3
			tempCols = append(tempCols, text.NewCol(1, fmt.Sprintf("%.1f", weatherData.Hourly.Temperature2m[index])))

			parsedTime, err := time.Parse("2006-01-02T15:04", weatherData.Hourly.Time[index])
			if err != nil {
				log.Fatal("Error parsing time")
			}

			timeCols = append(timeCols, text.NewCol(1, parsedTime.Format("15:04")))
		}

		parsedTime, err := time.Parse("2006-01-02T15:04", weatherData.Hourly.Time[i*24])
		if err != nil {
			log.Fatal("Error parsing time")
		}

		formattedTime := parsedTime.Format("January 02 Monday")

		m.AddRow(10, text.NewCol(5, formattedTime))
		m.AddRow(5, timeCols...)
		m.AddRow(10, tempCols...)
	}

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err)
	}

	err = document.Save(filename)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
