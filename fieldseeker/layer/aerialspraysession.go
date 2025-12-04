package layer

import (
	"encoding/json"
)

type AerialSpraySession struct {
	Geometry json.RawMessage
}

func (x *AerialSpraySession) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *AerialSpraySession) SetGeometry(m json.RawMessage) { x.Geometry = m }
