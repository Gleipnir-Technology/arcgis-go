package layer

import (
	"encoding/json"
)

type AerialSprayLine struct {
	Geometry json.RawMessage
}

func (x *AerialSprayLine) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *AerialSprayLine) SetGeometry(m json.RawMessage) { x.Geometry = m }
