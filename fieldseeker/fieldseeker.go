package fieldseeker

import (
	"errors"
	"github.com/Gleipnir-Technology/arcgis-go"
)

type FieldSeeker struct {
	Arcgis        *arcgis.ArcGIS
	FeatureServer *arcgis.FeatureServer
	ServiceInfo   *arcgis.ServiceInfo
	ServiceName   string
}

func NewFieldSeeker(ag *arcgis.ArcGIS, service string) *FieldSeeker {
	fs := new(FieldSeeker)
	fs.Arcgis = ag
	fs.FeatureServer = nil
	fs.ServiceName = service
	fs.ServiceInfo = nil
	return fs
}

// Make sure we have the Layer IDs we need to perform queries
func (fs *FieldSeeker) EnsureHasServiceInfo() error {
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

func stringOrEmpty(data map[string]any, key string) string {
	source, ok := data[key].(string)
	if ok {
		return source
	}
	return ""
}
