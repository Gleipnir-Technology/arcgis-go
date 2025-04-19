package main

import "flag"
import "fmt"

func main() {
	service_root := flag.String("root", "https://machine.domain.com", "The root URL for the ArcGIS service to contact.")
	tenant_id := flag.String("tenant", "abc123", "The tenant ID to use when communicating with the service.")
	token := flag.String("token", "secretstuff", "The API token to use to communicate with the service.")
	flag.Parse()

	arcgis := ArcGIS{*service_root, *tenant_id, *token}
	info, err := arcgis.info()
	if err != nil {
		fmt.Println("Failed: ", err)
	}
	fmt.Println("Current version: ", info.CurrentVersion)

	serviceInfo, err := arcgis.services()
	if err != nil {
		fmt.Println("Failed: ", err)
	}
	for _, service := range serviceInfo.Services {
		fmt.Println(service.Name)
	}
}
