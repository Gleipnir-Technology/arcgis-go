package layer

import (
	"encoding/json"
)

type QALarvCount struct {
	Geometry json.RawMessage
}

func (x *QALarvCount) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *QALarvCount) SetGeometry(m json.RawMessage) { x.Geometry = m }
