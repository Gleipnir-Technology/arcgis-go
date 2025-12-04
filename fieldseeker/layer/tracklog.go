package layer

import (
	"encoding/json"
)

type Tracklog struct {
	Geometry json.RawMessage
}

func (x *Tracklog) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *Tracklog) SetGeometry(m json.RawMessage) { x.Geometry = m }
