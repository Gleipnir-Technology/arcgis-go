package main

import (
	"fmt"
	"os"

	"github.com/Gleipnir-Technology/arcgis-go/examples"
)

func main() {
	fs, err := examples.FieldSeekerFromFlags()
	if err != nil {
		fmt.Println("Failed to create FS:", err)
		os.Exit(1)
	}

	t, err := fs.TrapData("2")
	if err != nil {
		fmt.Println("Failed: ", err)
		os.Exit(2)
	}
	if t != nil {
		fmt.Println("Trap Data: ", t)
	} else {
		fmt.Println("Nil trap")
	}
}
