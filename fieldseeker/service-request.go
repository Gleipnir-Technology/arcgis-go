package fieldseeker

import (
	"errors"
	"log"

	"github.com/Gleipnir-Technology/arcgis-go"
)

type FieldSeeker struct {
	Arcgis        *arcgis.ArcGIS
	FeatureServer *arcgis.FeatureServer
	ServiceInfo   *arcgis.ServiceInfo
	ServiceName   string
}

type ServiceRequest struct {
	Source string
}

func NewFieldSeeker(ag *arcgis.ArcGIS, service string) *FieldSeeker {
	fs := new(FieldSeeker)
	fs.Arcgis = ag
	fs.FeatureServer = nil
	fs.ServiceName = service
	fs.ServiceInfo = nil
	return fs
}

func (fs *FieldSeeker) ServiceRequest() (*ServiceRequest, error) {
	err := fs.ensureHasServiceInfo()
	if err != nil {
		return nil, err
	}
	var layer *arcgis.Layer
	for _, l := range fs.FeatureServer.Layers {
		if l.Name == "ServiceRequest" {
			layer = &l
		}
	}
	if layer == nil {
		return nil, errors.New("Unable to find ServiceRequest")
	}
	query := arcgis.NewQuery()
	query.Where = "1=1"
	query.OutFields = "*"
	results, err := fs.Arcgis.Query(
		fs.ServiceName,
		layer.ID,
		query,
	)
	log.Println("Results", results)
	if err != nil {
		return nil, err
	}
	if len(results.Features) == 0 {
		return nil, errors.New("Got no results")
	}
	f := results.Features[0]
	sr := new(ServiceRequest)
	source, ok := f.Attributes["SOURCE"].(string)
	if ok {
		sr.Source = source
	} else {
		return nil, errors.New("SOURCE not a string")
	}
	return sr, nil
}

// Make sure we have the Layer IDs we need to perform queries
func (fs *FieldSeeker) ensureHasServiceInfo() error {
	fs.ensureHasServices()
	if fs.ServiceInfo != nil {
		return nil
	}
	s, err := fs.Arcgis.FeatureServer(fs.ServiceName)
	if err != nil {
		return err
	}
	if s == nil {
		return errors.New("OH NOES")
	}
	fs.FeatureServer = s
	return nil
}

// Make sure we have the Service IDs we need to use FieldSeeker
func (fs FieldSeeker) ensureHasServices() error {
	if fs.ServiceInfo != nil {
		return nil
	}
	s, err := fs.Arcgis.Services()
	if err != nil {
		return err
	}
	fs.ServiceInfo = s
	return nil
}
