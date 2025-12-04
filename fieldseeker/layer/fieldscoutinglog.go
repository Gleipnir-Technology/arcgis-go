package layer

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type FieldScoutingLog struct {
	ObjectID       uint      `field:"OBJECTID"`
	Status         int16     `field:"STATUS"`
	GlobalID       uuid.UUID `field:"GlobalID"`
	CreatedUser    string    `field:"created_user"`
	CreatedDate    time.Time `field:"created_date"`
	LastEditedUser string    `field:"last_edited_user"`
	LastEditedDate time.Time `field:"last_edited_date"`
	CreationDate   time.Time `field:"CreationDate"`
	Creator        string    `field:"Creator"`
	EditDate       time.Time `field:"EditDate"`
	Editor         string    `field:"Editor"`
	Geometry       json.RawMessage
}

func (x *FieldScoutingLog) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *FieldScoutingLog) SetGeometry(m json.RawMessage) { x.Geometry = m }
