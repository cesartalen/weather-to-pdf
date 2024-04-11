package pdf

import (
	"cesartalen/weather-to-pdf/weather"
	"log"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
)

func GeneratePDF(filename string, weatherData weather.WeatherResponse) error {
	m := maroto.New()

	m.AddRow(20,
		text.NewCol(4, "Text"),
	)

	m.AddRow(10, col.New(12))

	m.AddRow(20,
		text.NewCol(4, "Text"),
	)

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
