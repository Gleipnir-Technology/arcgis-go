package fieldseeker

import (
	"errors"
	"fmt"
	"log"

	"github.com/Gleipnir-Technology/arcgis-go"
)

type ServiceRequest struct {
	Address          string
	City             string
	Description      string
	FieldNotes       string
	Location         arcgis.Geometry
	NotesForCustomer string
	NotesForTech     string
	Permission       string
	Priority         string
	Source           string
	Target           string
	Zip              string
}

func (fs *FieldSeeker) ServiceRequest() (*ServiceRequest, error) {
	err := fs.EnsureHasServiceInfo()
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
	if results.SpatialReference.WKID != 102100 {
		// See https://developers.arcgis.com/rest/services-reference/enterprise/geometry-objects/
		// for understanding how to support new spatial references
		return nil, fmt.Errorf("Unrecognized spatial reference %v", results.SpatialReference.WKID)
	}
	if len(results.Features) == 0 {
		return nil, errors.New("Got no results")
	}
	f := results.Features[0]
	sr := new(ServiceRequest)
	sr.Address = stringOrEmpty(f.Attributes, "REQADDR1") + stringOrEmpty(f.Attributes, "REQADDR2")
	sr.City = stringOrEmpty(f.Attributes, "REQCITY")
	sr.Description = stringOrEmpty(f.Attributes, "REQDESCR")
	sr.FieldNotes = stringOrEmpty(f.Attributes, "REQFLDNOTES")
	sr.Location = f.Geometry
	sr.NotesForCustomer = stringOrEmpty(f.Attributes, "REQNOTESFORCUST")
	sr.NotesForTech = stringOrEmpty(f.Attributes, "REQNOTESFORTECH")
	sr.Permission = stringOrEmpty(f.Attributes, "REQPERMISSION")
	sr.Priority = stringOrEmpty(f.Attributes, "PRIORITY")
	sr.Source = stringOrEmpty(f.Attributes, "SOURCE")
	sr.Target = stringOrEmpty(f.Attributes, "REQTARGET")
	sr.Zip = stringOrEmpty(f.Attributes, "REQZIP")
	return sr, nil
}
