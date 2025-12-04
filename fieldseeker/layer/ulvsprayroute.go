package layer

import (
	"encoding/json"
)

type ULVSprayRoute struct {
	Geometry json.RawMessage
}

func (x *ULVSprayRoute) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *ULVSprayRoute) SetGeometry(m json.RawMessage) { x.Geometry = m }
