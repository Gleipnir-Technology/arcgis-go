package layer

import (
	"encoding/json"
)

type OfflineMapAreas struct {
	Geometry json.RawMessage
}

func (x *OfflineMapAreas) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *OfflineMapAreas) SetGeometry(m json.RawMessage) { x.Geometry = m }
