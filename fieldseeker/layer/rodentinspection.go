package layer

import (
	"encoding/json"
)

type RodentInspection struct {
	Geometry json.RawMessage
}

func (x *RodentInspection) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *RodentInspection) SetGeometry(m json.RawMessage) { x.Geometry = m }
