package layer

import (
	"encoding/json"
)

type BarrierSpray struct {
	Geometry json.RawMessage
}

func (x *BarrierSpray) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *BarrierSpray) SetGeometry(m json.RawMessage) { x.Geometry = m }
