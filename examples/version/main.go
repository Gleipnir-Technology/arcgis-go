package main

import (
	"fmt"

	"github.com/Gleipnir-Technology/arcgis-go/examples"
)

func main() {
	ag, err := examples.ArcGISFromFlags()
	info, err := ag.Info()
	if err != nil {
		fmt.Println("Failed: ", err)
	}
	fmt.Println("Current version: ", info.CurrentVersion)
}
