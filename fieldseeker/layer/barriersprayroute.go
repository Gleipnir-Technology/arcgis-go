package layer

import (
	"encoding/json"
)

type BarrierSprayRoute struct {
	Geometry json.RawMessage
}

func (x *BarrierSprayRoute) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *BarrierSprayRoute) SetGeometry(m json.RawMessage) { x.Geometry = m }
