package layer

import (
	"encoding/json"
)

type QAProductObservation struct {
	Geometry json.RawMessage
}

func (x *QAProductObservation) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *QAProductObservation) SetGeometry(m json.RawMessage) { x.Geometry = m }
