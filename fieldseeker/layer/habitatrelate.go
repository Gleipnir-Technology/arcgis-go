package layer

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type HabitatRelate struct {
	ObjectID       uint      `field:"OBJECTID"`
	ForeignID      uuid.UUID `field:"FOREIGN_ID"`
	GlobalID       uuid.UUID `field:"GlobalID"`
	CreatedUser    string    `field:"created_user"`
	CreatedDate    time.Time `field:"created_date"`
	LastEditedUser string    `field:"last_edited_user"`
	LastEditedDate time.Time `field:"last_edited_date"`
	HabitatType    string    `field:"HABITATTYPE"`
	CreationDate   time.Time `field:"CreationDate"`
	Creator        string    `field:"Creator"`
	EditDate       time.Time `field:"EditDate"`
	Editor         string    `field:"Editor"`
	Geometry       json.RawMessage
}

func (x *HabitatRelate) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *HabitatRelate) SetGeometry(m json.RawMessage) { x.Geometry = m }
