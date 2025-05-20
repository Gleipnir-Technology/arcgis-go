package fieldseeker

import (
	"errors"
	"fmt"
	"log"

	"github.com/Gleipnir-Technology/arcgis-go"
)

type FieldSeeker struct {
	FeatureServer *arcgis.FeatureServer
	ServiceInfo   *arcgis.ServiceInfo
	ServiceName   string
}

var fs *FieldSeeker

func DoQuery(layer int, query *arcgis.Query) (*arcgis.QueryResult, error) {
	return arcgis.DoQuery(fs.ServiceName, layer, query)
}
func DoQueryRaw(layer int, query *arcgis.Query) ([]byte, error) {
	return arcgis.DoQueryRaw(fs.ServiceName, layer, query)
}

func FeatureServerLayers() []arcgis.Layer {
	return fs.FeatureServer.Layers
}

func Initialize(
	service_root string,
	tenant_id string,
	token string,
	service_name string) error {
	arcgis.Initialize(
		service_root,
		tenant_id,
		token)

	fs = new(FieldSeeker)
	fs.FeatureServer = nil
	fs.ServiceName = service_name
	fs.ServiceInfo = nil

	err := fs.ensureHasServiceInfo()
	if err != nil {
		return fmt.Errorf("Failed to get FieldSeeker service info: %v", err)
	}
	log.Println("Connected to FieldSeeker")
	return nil
}

func MaxRecordCount() int {
	if fs == nil {
		return 0
	}
	return fs.FeatureServer.MaxRecordCount
}

func QueryCount(layer_id int) (*arcgis.QueryResultCount, error) {
	return arcgis.QueryCount(fs.ServiceName, layer_id)
}

// Make sure we have the Layer IDs we need to perform queries
func (fs *FieldSeeker) ensureHasServiceInfo() error {
	fs.ensureHasServices()
	if fs.ServiceInfo != nil {
		return nil
	}
	s, err := arcgis.GetFeatureServer(fs.ServiceName)
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
	s, err := arcgis.Services()
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
