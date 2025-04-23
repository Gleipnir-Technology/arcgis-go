package examples

import (
	"flag"

	"github.com/Gleipnir-Technology/arcgis-go"
	"github.com/Gleipnir-Technology/arcgis-go/fieldseeker"
)

func ArcGISFromFlags() (*arcgis.ArcGIS, error) {
	service_root := flag.String("root", "https://machine.domain.com", "The root URL for the ArcGIS service to contact.")
	tenant_id := flag.String("tenant", "abc123", "The tenant ID to use when communicating with the service.")
	token := flag.String("token", "secretstuff", "The API token to use to communicate with the service.")
	flag.Parse()

	arcgis := arcgis.ArcGIS{*service_root, *tenant_id, *token}
	return &arcgis, nil
}

func FieldSeekerFromFlags() (*fieldseeker.FieldSeeker, error) {
	service := flag.String("service", "some-service", "The service that holds FieldSeeker data")
	ag, err := ArcGISFromFlags()
	if err != nil {
		return nil, err
	}

	fs := fieldseeker.NewFieldSeeker(ag, *service)
	return fs, nil
}
