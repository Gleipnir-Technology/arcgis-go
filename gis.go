package main

import (
	"encoding/json"
)

type ServiceListing struct {
	Name string
	Type string
	URL string
}

type ServiceInfo struct {
	CurrentVersion float64
	Services[] ServiceListing
}

func ParseServiceInfo(data []byte) (*ServiceInfo, error) {
	var result ServiceInfo
	err := json.Unmarshal(data, &result);
	if err != nil {
		return nil, err
	}
	return &result, nil
}
