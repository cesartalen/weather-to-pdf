package pdf

import (
	"cesartalen/weather-to-pdf/weather"
	"fmt"
	"log"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
)

func GeneratePDF(filename string, weatherData weather.WeatherResponse) error {
	m := maroto.New()

	for i := 0; i < len(weatherData.Hourly.Time); i++ {
		if (i % 24) == 0 {
		}

		temp := fmt.Sprintf("%.1f", weatherData.Hourly.Temperature2m[i])

		m.AddRow(5,
			text.NewCol(4, string(weatherData.Hourly.Time[i])),
			text.NewCol(4, temp),
		)
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
