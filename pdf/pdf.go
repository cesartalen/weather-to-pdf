package pdf

import (
	"cesartalen/weather-to-pdf/weather"
	"fmt"
	"log"
	"time"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func GeneratePDF(filename string, weatherData weather.WeatherResponse) error {
	cfg := config.NewBuilder().WithMargins(10, 15, 10).Build()
	m := maroto.New(cfg)

	m.AddRow(5, text.NewCol(12, "Weather Forecast", props.Text{Align: align.Right}))
	m.AddRow(5, text.NewCol(12, time.Now().Format("2006-01-02 15:04"), props.Text{Align: align.Right}))
	m.AddRow(12, text.NewCol(12, fmt.Sprintf("Timezone: %s", weatherData.TimezoneAbbrev), props.Text{Align: align.Right}))
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
			tempCols = append(tempCols, text.NewCol(1, fmt.Sprintf("%.1f °C", weatherData.Hourly.Temperature2m[index]), props.Text{Left: 20}))

			parsedTime, err := time.Parse("2006-01-02T15:04", weatherData.Hourly.Time[index])
			if err != nil {
				log.Fatal("Error parsing time")
			}

			timeCols = append(timeCols, text.NewCol(1, parsedTime.Format("15:04"), props.Text{Left: 14}))
		}

		parsedTime, err := time.Parse("2006-01-02T15:04", weatherData.Hourly.Time[i*24])
		if err != nil {
			log.Fatal("Error parsing time")
		}

		formattedTime := parsedTime.Format("January 02 Monday")

		m.AddRow(5, text.NewCol(10, formattedTime, props.Text{
			Size: 12,
			Left: 10,
		}))
		m.AddRow(5, timeCols...)
		m.AddRow(20, tempCols...)
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
