package main

import (
	"encoding/json"
)

// Root structure for an instance of the ArcGIS API
type ArcGIS struct {
	ServiceRoot string
	TenantId string
	Token string
}

// Basic information about the REST API itself
type AuthInfo struct {
	isTokenBasedSecurity bool
	tokenServiceUrl string
}
type RestInfo struct {
	CurrentVersion float64
	FullVersion string
	OwningSystemUrl string
	OwningTenant string
	AuthInfo AuthInfo
}

// Listing of available services
type ServiceListing struct {
	Name string
	Type string
	URL string
}

type ServiceInfo struct {
	CurrentVersion float64
	Services[] ServiceListing
}

func ParseRestInfo(data []byte) (*RestInfo, error) {
	var result RestInfo
	err := json.Unmarshal(data, &result);
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func ParseServiceInfo(data []byte) (*ServiceInfo, error) {
	var result ServiceInfo
	err := json.Unmarshal(data, &result);
	if err != nil {
		return nil, err
	}
	return &result, nil
}
