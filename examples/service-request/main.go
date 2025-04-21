package main

import (
	"fmt"

	"github.com/Gleipnir-Technology/arcgis-go/examples"
)

func main() {
	ag, err := examples.ArcGISFromFlags()
	if err != nil {
		fmt.Println("Failed to create ARCGIS")
	}

	ag.Info()
}
