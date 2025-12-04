package layer

import (
	"encoding/json"
)

type LandingCount struct {
	Geometry json.RawMessage
}

func (x *LandingCount) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *LandingCount) SetGeometry(m json.RawMessage) { x.Geometry = m }
