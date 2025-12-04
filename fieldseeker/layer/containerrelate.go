package layer

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type ContainerRelate struct {
	ObjectID       uint      `field:"OBJECTID"`
	GlobalID       uuid.UUID `field:"GlobalID"`
	CreatedUser    string    `field:"created_user"`
	CreatedDate    time.Time `field:"created_date"`
	LastEditedUser string    `field:"last_edited_user"`
	LastEditedDate time.Time `field:"last_edited_date"`
	InspsampleID   uuid.UUID `field:"INSPSAMPLEID"`
	MosquitoinspID uuid.UUID `field:"MOSQUITOINSPID"`
	TreatmentID    uuid.UUID `field:"TREATMENTID"`
	ContainerType  string    `field:"CONTAINERTYPE"`
	CreationDate   time.Time `field:"CreationDate"`
	Creator        string    `field:"Creator"`
	EditDate       time.Time `field:"EditDate"`
	Editor         string    `field:"Editor"`
	Geometry       json.RawMessage
}

func (x *ContainerRelate) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *ContainerRelate) SetGeometry(m json.RawMessage) { x.Geometry = m }
