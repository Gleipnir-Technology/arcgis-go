package main

import (
	"io"
	"os"
	"testing"
)

func readFileOrFail(t *testing.T, filename string) ([]byte, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		t.Errorf(`Failed to open %s`, filename)
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	return byteValue, nil
}

func TestParseRestInfo(t *testing.T) {
	content, err := readFileOrFail(t, "rest.json")
	if err != nil {
		return
	}
	rest, err := ParseRestInfo(content)
	if err != nil {
		t.Error("Failed to parse")
	}
	if rest.CurrentVersion != 11.2 {
		t.Error("Incorrect version")
	}
}

func TestParseServices(t *testing.T) {
	content, err := readFileOrFail(t, "services.json")
	if err != nil {
		return
	}
	
	services, err := ParseServiceInfo(content)
	if err != nil {
		t.Error("Failed to parse")
	}
	if len(services.Services) != 55 {
		t.Errorf(`Wrong services. Expected 55, got %d`, len(services.Services))
	}
	if services.Services[0].Name != "_APR_AUG_2024_Aegypti_Abundance" {
		t.Error("Wrong service name")
	}
}
