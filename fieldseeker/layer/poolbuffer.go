package layer

import (
	"encoding/json"
)

type PoolBuffer struct {
	Geometry json.RawMessage
}

func (x *PoolBuffer) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *PoolBuffer) SetGeometry(m json.RawMessage) { x.Geometry = m }
