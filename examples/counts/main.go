package main

import (
	"fmt"
	"os"

	"github.com/Gleipnir-Technology/arcgis-go/examples"
)

// Get counts for every table and layer.
func main() {
	fs, err := examples.FieldSeekerFromFlags()
	if err != nil {
		fmt.Println("Failed to create FS:", err)
		os.Exit(1)
	}

	fs.EnsureHasServiceInfo()
	fmt.Println("Layers")
	for _, layer := range fs.FeatureServer.Layers {
		c, err := fs.Arcgis.QueryCount(fs.ServiceName, layer.ID)
		if err != nil {
			fmt.Println("Failed", err)
			os.Exit(1)
		}
		fmt.Printf("%v %v %v\n", layer.ID, c.Count, layer.Name)
	}
	fmt.Println("Tables")
	for _, table := range fs.FeatureServer.Tables {
		c, err := fs.Arcgis.QueryCount(fs.ServiceName, table.ID)
		if err != nil {
			fmt.Println("Failed", err)
			os.Exit(1)
		}
		fmt.Printf("%v %v %v\n", table.ID, c.Count, table.Name)
	}
}
