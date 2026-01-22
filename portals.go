package arcgis
import (
	"encoding/json"
)
func (ag *ArcGIS) PortalsSelf() (*PortalsResponse, error) {
	// We may need to always direct this request to
	// https://www.arcgis.com/sharing/rest/portals/self?f=json
	// not sure if hosted services are different
	req, err := ag.sharingRequest("/portals/self")
	if err != nil {
		return nil, err
	}
	content, err := ag.requestJSON(req)
	if err != nil {
		return nil, err
	}
	return parsePortalsResponse(content)
}

func parsePortalsResponse(data []byte) (*PortalsResponse, error) {
	var result PortalsResponse
	saveResponse(data, "portal.json")
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

