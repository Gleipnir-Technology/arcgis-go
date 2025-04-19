package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Root structure for an instance of the ArcGIS API
type ArcGIS struct {
	ServiceRoot string
	TenantId    string
	Token       string
}

// Basic information about the REST API itself
type AuthInfo struct {
	isTokenBasedSecurity bool
	tokenServiceUrl      string
}
type RestInfo struct {
	CurrentVersion  float64
	FullVersion     string
	OwningSystemUrl string
	OwningTenant    string
	AuthInfo        AuthInfo
}

// Listing of available services
type ServiceListing struct {
	Name string
	Type string
	URL  string
}

type ServiceInfo struct {
	CurrentVersion float64
	Services       []ServiceListing
}

func parseRestInfo(data []byte) (*RestInfo, error) {
	var result RestInfo
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func parseServiceInfo(data []byte) (*ServiceInfo, error) {
	var result ServiceInfo
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (arcgis ArcGIS) serviceUrl(endpoint string) (*url.URL, error) {
	u := fmt.Sprintf("%s/%s/arcgis/rest/%s", arcgis.ServiceRoot, arcgis.TenantId, endpoint)
	base, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("f", "json")
	params.Add("token", arcgis.Token)
	base.RawQuery = params.Encode()
	return base, nil
}

func requestJSON(u *url.URL) ([]byte, error) {
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (arcgis ArcGIS) info() (*RestInfo, error) {
	u, err := arcgis.serviceUrl("/info")
	if err != nil {
		return nil, err
	}
	content, err := requestJSON(u)
	if err != nil {
		return nil, err
	}
	return parseRestInfo(content)
}

func (arcgis ArcGIS) services() (*ServiceInfo, error) {
	u, err := arcgis.serviceUrl("/services")
	if err != nil {
		return nil, err
	}
	content, err := requestJSON(u)
	if err != nil {
		return nil, err
	}
	return parseServiceInfo(content)
}
