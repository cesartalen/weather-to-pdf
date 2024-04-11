package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg/core"

	"github.com/johnfercher/maroto/v2/pkg/components/col"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/text"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err)
	}

	err = document.Save("weather.pdf")
	if err != nil {
		log.Fatal(err)
	}
}

func GetMaroto() core.Maroto {
	m := maroto.New()

	m.AddRow(20,
		text.NewCol(4, "Text"),
	)

	m.AddRow(10, col.New(12))

	m.AddRow(20,
		text.NewCol(4, "Text"),
	)

	return m
}
