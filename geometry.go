package arcgis

import (
	"context"

	"github.com/Gleipnir-Technology/arcgis-go/response"
)

// This isn't tested
func (ag *ArcGIS) GeometryProject(ctx context.Context, geometries []response.Geometry) (*response.Project, error) {
	params := make(map[string]string)
	params["f"] = "json"
	params["geometries"] = ""
	params["inSR"] = ""
	params["outSR"] = ""
	url := ag.urlGeometry.JoinPath("project")
	return reqPostFormToJSONFullURL[response.Project](ctx, ag.requestor, *url, params)
}
