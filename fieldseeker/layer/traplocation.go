package layer

import (
	"github.com/Gleipnir-Technology/arcgis-go/response"
	"time"

	"github.com/google/uuid"
)

type TrapLocation struct {
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
	Geometry       response.Geometry
}

func (x *TrapLocation) GetGeometry() response.Geometry  { return x.Geometry }
func (x *TrapLocation) SetGeometry(m response.Geometry) { x.Geometry = m }
