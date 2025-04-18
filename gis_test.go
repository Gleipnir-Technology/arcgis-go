package main

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestParseServices(t *testing.T) {
	jsonFile, err := os.Open("services.json")
	if err != nil {
		t.Errorf("Failed to open JSON data")
	}
	fmt.Println("Opened services.json")
	byteValue, _ := io.ReadAll(jsonFile)
	
	services, err := ParseServiceInfo(byteValue)
	if err != nil {
		t.Error("Failed to parse")
	}
	if len(services.Services) != 55 {
		t.Errorf(`Wrong services. Expected 55, got %d`, len(services.Services))
	}
	if services.Services[0].Name != "_APR_AUG_2024_Aegypti_Abundance" {
		t.Error("Wrong service name")
	}
	defer jsonFile.Close()
}
