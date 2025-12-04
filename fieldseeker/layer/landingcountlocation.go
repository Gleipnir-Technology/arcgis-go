package layer

import (
	"encoding/json"
)

type LandingCountLocation struct {
	Geometry json.RawMessage
}

func (x *LandingCountLocation) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *LandingCountLocation) SetGeometry(m json.RawMessage) { x.Geometry = m }
