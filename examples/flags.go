package examples

import (
	"flag"

	"github.com/Gleipnir-Technology/arcgis-go"
)

func ArcGISFromFlags() (*arcgis.ArcGIS, error) {
	service_root := flag.String("root", "https://machine.domain.com", "The root URL for the ArcGIS service to contact.")
	tenant_id := flag.String("tenant", "abc123", "The tenant ID to use when communicating with the service.")
	token := flag.String("token", "secretstuff", "The API token to use to communicate with the service.")
	flag.Parse()

	arcgis := arcgis.ArcGIS{*service_root, *tenant_id, *token}
	return &arcgis, nil
}
