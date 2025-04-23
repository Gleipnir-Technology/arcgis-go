package main

import (
	"fmt"
	"os"

	"github.com/Gleipnir-Technology/arcgis-go"
	"github.com/Gleipnir-Technology/arcgis-go/examples"
)

func main() {
	fs, err := examples.FieldSeekerFromFlags()
	if err != nil {
		fmt.Println("Failed to create FS:", err)
		os.Exit(1)
	}
	sr, err := fs.ServiceRequest()
	if err != nil {
		fmt.Println("Failed: ", err)
		os.Exit(2)
	}
	if sr != nil {
		fmt.Println("Field Request: ", sr)
		fmt.Println("Lat/Long: ", arcgis.GeometryToPosition(sr.Location))
	} else {
		fmt.Println("Nil service request")
	}
}
