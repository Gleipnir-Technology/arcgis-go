package arcgis

import (
	"context"
)

func (ag *ArcGIS) PortalsSelf(ctx context.Context) (*PortalsResponse, error) {
	// We may need to always direct this request to
	// https://www.arcgis.com/sharing/rest/portals/self?f=json
	// not sure if hosted services are different
	return doJSONGet[PortalsResponse](ctx, ag.requestor, "/portals/self")
}
