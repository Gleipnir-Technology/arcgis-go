package layer

import (
	"encoding/json"
)

type RestrictedArea struct {
	Geometry json.RawMessage
}

func (x *RestrictedArea) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *RestrictedArea) SetGeometry(m json.RawMessage) { x.Geometry = m }
