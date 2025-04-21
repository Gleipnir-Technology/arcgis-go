package main

import "fmt"

func main() {
	ag, err := ArcGISFromFlags()
	if err != nil {
		fmt.Println("Failed to create ARCGIS")
	}

	ag.info()
}
