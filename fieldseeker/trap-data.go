package fieldseeker

import (
	"errors"
	"fmt"
	//"log"

	"github.com/Gleipnir-Technology/arcgis-go"
)

type TrapLocation struct {
	AccessDescription string
	Comments          string
	Description       string
	Name              string
}

type TrapData struct {
	Comments  string
	Condition string
	End       string
	FieldTech string
	Location  TrapLocation
	Type      string
}

func (fs *FieldSeeker) TrapData(objectID string) (*TrapData, error) {
	err := fs.ensureHasServiceInfo()
	if err != nil {
		return nil, err
	}
	var layer_location *arcgis.Layer
	var layer_data *arcgis.Layer
	for _, l := range fs.FeatureServer.Layers {
		if l.Name == "TrapLocation" {
			layer_location = &l
		}
		if l.Name == "TrapData" {
			layer_data = &l
		}
	}
	if layer_location == nil {
		return nil, errors.New("Unable to find TrapLocation")
	}
	if layer_data == nil {
		return nil, errors.New("Unable to find TrapData")
	}

	query := arcgis.NewQuery()
	query.Limit = 10
	query.Where = "1=1"
	//query.ObjectIDs = objectID
	query.OutFields = "*"
	results, err := fs.Arcgis.Query(
		fs.ServiceName,
		layer_data.ID,
		query,
	)
	//log.Println("Layer Data Results", results)
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
	td := new(TrapData)
	td.Comments = stringOrEmpty(f.Attributes, "COMMENTS")
	td.Condition = stringOrEmpty(f.Attributes, "TRAPCONDITION")
	td.End = stringOrEmpty(f.Attributes, "ENDATETIME")
	td.FieldTech = stringOrEmpty(f.Attributes, "FIELDTECH")
	td.Type = stringOrEmpty(f.Attributes, "TRAPTYPE")

	location_id := stringOrEmpty(f.Attributes, "LOC_ID")
	if len(location_id) == 0 {
		return nil, errors.New("No LOC_ID")
	}

	query = arcgis.NewQuery()
	query.Limit = 10
	query.Where = fmt.Sprintf("GlobalID='%v'", location_id)
	query.OutFields = "*"
	results, err = fs.Arcgis.Query(
		fs.ServiceName,
		layer_location.ID,
		query,
	)
	if err != nil {
		return nil, err
	}
	f = results.Features[0]
	td.Location.AccessDescription = stringOrEmpty(f.Attributes, "ACCESSDESC")
	td.Location.Comments = stringOrEmpty(f.Attributes, "COMMENTS")
	td.Location.Description = stringOrEmpty(f.Attributes, "DESCRIPTION")
	td.Location.Name = stringOrEmpty(f.Attributes, "NAME")

	//log.Println("Layer Location Results", results)

	// Now we have *some* of the data. Get more of it.
	return td, nil
}
